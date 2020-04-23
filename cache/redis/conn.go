package redis

import (
	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
	redisHost="47.111.181.52:6379"
	redisPass="rechengparty"
)

func 