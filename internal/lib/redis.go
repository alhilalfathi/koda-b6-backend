package lib

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "server-hilal-redis-1:6379",
	})
}

var Ctx = context.Background()
