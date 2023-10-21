package stores

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/m-posluszny/go-ynab/src/config"
)

func GetRedisSessionStore(authConf config.AuthConf, redisConf config.RedisConf) redis.Store {
	store, err := redis.NewStore(redisConf.Size, "tcp", redisConf.Host, redisConf.Password, []byte(authConf.Secret))
	if err != nil {
		panic(err)
	}
	return store
}
