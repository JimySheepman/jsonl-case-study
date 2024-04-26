//go:generate mockgen -package=mocks -destination=../../../../mocks/redis_mock.go -source=redis.go

package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type redisImpl struct {
	redisC redisClient
}

func NewRedisClient(r redisClient) *redisImpl {
	return &redisImpl{
		redisC: r,
	}
}

func (c *redisImpl) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.redisC.Get(ctx, key)
}

func (c *redisImpl) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.redisC.Set(ctx, key, value, expiration)
}

func (c *redisImpl) GetBody(ctx context.Context, key string, parser interface{}) (bool, error) {
	result, err := c.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, err
	}

	if err = json.Unmarshal([]byte(result), parser); err != nil {
		return false, err
	}

	return true, nil
}
