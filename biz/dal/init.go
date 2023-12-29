package dal

import (
	"github.com/wheelergeo/g-otter-gateway/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
