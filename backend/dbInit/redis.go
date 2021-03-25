package dbInit

import (
	"github.com/go-redis/redis/v8"
	config "yxcblog/config"
)

var REDIS *redis.Client

func init(){
	REDIS = redis.NewClient(&redis.Options{
		Addr:     config.IPADDRESS+":"+config.REDIS_PORT,
		Password: config.REDIS_PASSWORD, // no password set,
		DB: 0,
	})
}
