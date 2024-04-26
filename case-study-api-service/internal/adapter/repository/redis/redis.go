package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
)

type redisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
}

type redisImpl struct {
	keyPrefix string
	redisC    redisClient
}

func NewRedisClient(r redisClient) *redisImpl {
	return &redisImpl{
		redisC: r,
	}
}

func (c *redisImpl) GetProduct(ctx context.Context, key string, parser interface{}) (bool, error) {
	result, err := c.redisC.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, err
	}

	if err := json.Unmarshal([]byte(result), parser); err != nil {
		return false, err
	}

	return true, nil
}
