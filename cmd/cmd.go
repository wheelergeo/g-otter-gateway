package gatewaycmd

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/paseto"
	"github.com/hertz-contrib/pprof"
	"github.com/spf13/cobra"
	"github.com/wheelergeo/g-otter-gateway/biz/dal"
	"github.com/wheelergeo/g-otter-gateway/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-gateway/biz/router"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	"github.com/wheelergeo/g-otter-gateway/conf"
	"github.com/wheelergeo/g-otter-gateway/pkg/auth"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var c *cobra.Command

func Command() *cobra.Command {
	once.Do(func() {
		c = &cobra.Command{
			Use:   "gateway",
			Short: "start gateway",
			Run: func(cmd *cobra.Command, args []string) {
				dal.Init()
				rpc.Init()
				auth.NewServer(
					conf.GetConf().Casbin.ModelName,
					mysql.DB,
					conf.GetConf().Casbin.PolicyTable,
					conf.GetConf().Casbin.PolicyRedis,
				)

				h := server.New(
					server.WithHostPorts(conf.GetConf().Hertz.Address),
				)

				registerMiddleware(h)
				router.GeneratedRegister(h)

				h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
					ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
				})
				h.GET("/paseto", func(c context.Context, ctx *app.RequestContext) {
					now := time.Now()
					genTokenFunc := paseto.DefaultGenTokenFunc()
					token, err := genTokenFunc(&paseto.StandardClaims{
						Issuer:    "cwg-issuer",
						ExpiredAt: now.Add(time.Hour),
						NotBefore: now,
						IssuedAt:  now,
					}, nil, nil)
					if err != nil {
						hlog.Error("generate token failed")
					}
					ctx.String(http.StatusOK, token)

				})

				h.POST("/paseto", paseto.New(), func(c context.Context, ctx *app.RequestContext) {
					ctx.String(http.StatusOK, "token is valid")

				})

				h.Spin()
			},
		}
	})

	return c
}

func registerMiddleware(h *server.Hertz) {
	// log
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}

	hlog.SetOutput(io.MultiWriter(asyncWriter, os.Stdout))

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
