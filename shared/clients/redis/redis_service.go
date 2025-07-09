package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	rdb *redis.Client
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.rdb.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r *redisClient) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}

func (r *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	val, err := r.rdb.Exists(ctx, key).Result()
	return val > 0, err
}

func (r *redisClient) Close() error {
	return r.rdb.Close()
}
