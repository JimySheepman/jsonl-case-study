//go:build integration

package integrations

import (
	repository "case-study-job-service/internal/adapter/repository/redis"
	"case-study-job-service/internal/core/port"
	redisstorage "case-study-job-service/internal/platform/storage/redis"
	"case-study-job-service/pkg/config"
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"log"
	"os"
	"testing"
)

var redisRepo port.ProductRepository

func TestMain(m *testing.M) {
	ctx := context.Background()

	redisContainer, err := redis.RunContainer(ctx,
		testcontainers.WithImage("docker.io/redis:7"),
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	state, err := redisContainer.State(ctx)
	if err != nil {
		log.Fatalf("failed to get container state: %s", err) // nolint:gocritic
	}

	fmt.Println(state.Running)

	redisHost, err := redisContainer.Host(ctx)
	if err != nil {
		log.Fatalf("%s get container host failed", err.Error())
	}

	redisPort, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		log.Fatalf("%s mapped container port failed", err.Error())
	}

	if err := config.LoadConfig("../../"); err != nil {
		log.Fatalf("%s config load failed", err.Error())
	}

	config.Cfg.Redis.Address = redisHost + ":" + redisPort.Port()

	redisConn := redisstorage.RedisConnection()
	redisRepo = repository.NewRedisClient(redisConn)

	os.Exit(m.Run())
}
