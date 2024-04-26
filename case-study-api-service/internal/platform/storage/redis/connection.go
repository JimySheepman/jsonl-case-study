package storage

import (
	"case-study-api-service/pkg/config"
	"case-study-api-service/pkg/logger"
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	client redis.UniversalClient
	once   sync.Once
)

func RedisConnection() redis.UniversalClient {
	l := logger.GetLogger().Sugar()

	once.Do(func() {
		l.Info("Redis connection address: ", config.Cfg.Redis.Address)

		o := &redis.UniversalOptions{
			Addrs:           []string{config.Cfg.Redis.Address},
			Username:        config.Cfg.Redis.Username,
			Password:        config.Cfg.Redis.Password,
			MaxRedirects:    8,
			MaxRetries:      50,
			MinRetryBackoff: time.Millisecond * 10,
			MaxRetryBackoff: time.Millisecond * 200,
		}

		client = redis.NewUniversalClient(o)

		l.Info("trying to ping redis server")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if _, err := client.Ping(ctx).Result(); err != nil {
			panic(err)
		}

		l.Info("redis server ping and connection successful")
	})

	return client
}
