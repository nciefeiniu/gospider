package public

import (
	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func init() {
	host := "127.0.0.1:6379"
	db := redis.DialDatabase(5)
	Pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host, db)
		},
	}
}

func Get() redis.Conn {
	return Pool.Get()
}