package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

type redisCache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) Cache {
	return &redisCache{rdb: rdb}
}

// Get
func (c *redisCache) Get(ctx context.Context, key string, dest any) error {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// Set
func (c *redisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key, data, ttl).Err()
}

// Delete
func (c *redisCache) Delete(ctx context.Context, keys ...string) error {
	return c.rdb.Del(ctx, keys...).Err()
}
