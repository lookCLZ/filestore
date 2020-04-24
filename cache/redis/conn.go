package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool      *redis.Pool
	redisHost = "47.111.181.52:6379"
	redisPass = "rechengparty"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			// 访问认证
			if _, err = c.Do("AUTH", redisPass); err != nil {
				c.Close()
			}
			return c, nil
		},
		// TestOnBorrow is an optional application supplied function for checking
		// the health of an idle connection before the connection is used again by
		// the application. Argument t is the time that the connection was returned
		// to the pool. If the function returns an error, then the connection is
		// closed.
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			// Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
