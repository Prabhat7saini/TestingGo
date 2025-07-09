package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
)

var (
	instance Client
	once     sync.Once
)

func InitRedis(cfg *config.Env) (Client, error) {
	var err error

	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.Db,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if pingErr := rdb.Ping(ctx).Err(); pingErr != nil {
			err = fmt.Errorf("failed to connect to Redis at %s: %w", cfg.Redis.Addr, pingErr)
			return
		}

		instance = &redisClient{rdb: rdb}
	})

	if instance == nil {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("redis client is nil (unknown error)")
	}

	return instance, nil
}

func GetRedisClient() (Client, error) {
	if instance == nil {
		return nil, fmt.Errorf("redis client not initialized ")
	}
	return instance, nil
}
