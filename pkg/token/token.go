package token

import (
	"context"
	"encoding"
	"encoding/json"
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/redis/go-redis/v9"
)

var _ encoding.BinaryMarshaler = new(TokenValue)
var _ encoding.BinaryUnmarshaler = new(TokenValue)

type STokenAuth struct {
	redis    *redis.Client
	cacheKey string
	parser   paseto.Parser
}

type ClaimData map[string]interface{}

type TokenValue struct {
	Key           [32]byte  `json:"key"`
	Authorization string    `json:"authorization"`
	Refresh       int       `json:"refresh"`
	Timeout       int       `json:"timeout"`
	IssuedAt      time.Time `json:"issueAt"`
}

var sTokenAuth *STokenAuth

func Init(redis *redis.Client, cacheKey string) {
	sTokenAuth = &STokenAuth{
		redis:    redis,
		cacheKey: cacheKey,
		parser:   paseto.NewParser(),
	}
}

func TokenAuth() *STokenAuth {
	return sTokenAuth
}

func (m TokenValue) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *TokenValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (ta *STokenAuth) New(refresh, timeout int,
	data ClaimData) (authorization string, err error) {

	duration := time.Duration(timeout * int(time.Second))
	tokenValue, _, err := ta.newToken(refresh, timeout, data)
	if err != nil {
		return
	}

	authorization = tokenValue.Authorization
	err = ta.redis.Set(
		context.Background(),
		ta.cacheKey+":"+tokenValue.Authorization,
		*tokenValue,
		duration,
	).Err()

	return
}

func (ta *STokenAuth) newToken(refresh, timeout int,
	data ClaimData) (tokenValue *TokenValue, pToken paseto.Token, err error) {

	t := time.Now()
	key := paseto.NewV4SymmetricKey()
	duration := time.Duration(timeout * int(time.Second))
	tokenValue = &TokenValue{
		Refresh:  refresh,
		Timeout:  timeout,
		IssuedAt: t,
		Key:      [32]byte(key.ExportBytes()),
	}

	pToken = paseto.NewToken()
	pToken.SetIssuedAt(t)
	pToken.SetNotBefore(t)
	pToken.SetExpiration(t.Add(duration))

	for k, v := range data {
		if err = pToken.Set(k, v); err != nil {
			return
		}
	}

	tokenValue.Authorization = pToken.V4Encrypt(key, nil)

	return
}

func (ta *STokenAuth) Parse(authorization string) (data ClaimData, err error) {
	// 1. get token from redis
	cmd := ta.redis.Get(context.Background(), ta.cacheKey+":"+authorization)
	if cmd.Err() == redis.Nil {
		err = errors.New("Token is expired")
		return
	}

	tokenValue := &TokenValue{}
	err = cmd.Scan(tokenValue)
	if err != nil {
		return
	}

	// 2. parse token
	key, _ := paseto.V4SymmetricKeyFromBytes(tokenValue.Key[:])
	pToken, err := ta.parser.ParseV4Local(
		key,
		tokenValue.Authorization,
		nil,
	)
	if err != nil {
		return
	}
	data = pToken.Claims()

	// 3. judge token is need to be refresh or not
	t := time.Now()
	tRefresh := tokenValue.IssuedAt.Add(time.Second * time.Duration(tokenValue.Refresh))
	tTimeout := tokenValue.IssuedAt.Add(time.Second * time.Duration(tokenValue.Timeout))
	if t.After(tRefresh) && t.Before(tTimeout) {
		_, err = ta.refresh(tokenValue, pToken)
		if err != nil {
			return
		}
		return
	}
	return
}

func (ta *STokenAuth) refresh(oldTokenValue *TokenValue,
	oldPToken *paseto.Token) (newPToken paseto.Token, err error) {

	duration := time.Duration(oldTokenValue.Timeout * int(time.Second))

	// 1. gen new token
	newTokenValue, newPToken, err := ta.newToken(
		oldTokenValue.Refresh,
		oldTokenValue.Timeout,
		oldPToken.Claims(),
	)
	if err != nil {
		return
	}

	// 2. replace token
	err = ta.redis.Set(
		context.Background(),
		ta.cacheKey+":"+oldTokenValue.Authorization,
		*newTokenValue,
		duration,
	).Err()
	return
}

func (ta *STokenAuth) Delete(authorization string) (err error) {
	return ta.redis.Del(context.Background(), ta.cacheKey+":"+authorization).Err()
}
