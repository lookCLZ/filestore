package redis

import (
	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
	redisHost="47.111.181.52:6379"
	redisPass="rechengparty"
)

func newRedisPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle:50,
		MaxActive: 30,
		IdleTimeout: 300 * time.Second,
		Dial:func()(redis.Conn,error){
			c,err:=reids
		}
	}
}