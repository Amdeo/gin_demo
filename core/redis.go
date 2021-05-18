package core

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"server/global"
)

func Redis() {
	redisCfg := global.GConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		zap.L().Info("redis connect ping response:", zap.String("pong", pong))
		global.GRedis = client
	}
}
