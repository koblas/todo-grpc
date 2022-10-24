package redisutil

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func NewClient(addr string) *redis.Client {
	if addr == "" {
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "",                     // no password set
		DB:          0,                      // use default DB
		DialTimeout: time.Millisecond * 200, // either it happens or it doesn't
	})

	return rdb
}
