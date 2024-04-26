package main

import (
	handler "case-study-job-service/internal/adapter/handler/job"
	repository "case-study-job-service/internal/adapter/repository/redis"
	s3repository "case-study-job-service/internal/adapter/repository/s3"
	service "case-study-job-service/internal/core/service/job"
	server "case-study-job-service/internal/platform/server/job"
	redisstorage "case-study-job-service/internal/platform/storage/redis"
	s3storage "case-study-job-service/internal/platform/storage/s3"
	"case-study-job-service/pkg/config"
	"case-study-job-service/pkg/logger"
	"context"
	"log"
	"runtime"
	"time"
)

const worker = 3

func main() {
	now := time.Now()
	defer func() {
		log.Printf("case-study-job-service is finished, latency: %s", time.Since(now))
	}()

	if err := config.LoadConfig("."); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	l := logger.RegisterLogger(config.Cfg.LoggerConfig)
	ctx := logger.WithCtx(context.TODO(), l)
	lg := l.Sugar()

	lg.Info("case-study-job-service is started...")
	lg.Infof("Runtime config %+v", config.Cfg)
	lg.Infof("Go runtime is %s", runtime.Version())

	redisConn := redisstorage.RedisConnection()
	s3Conn := s3storage.S3Connection()

	redisClient := repository.NewRedisClient(redisConn)
	s3Client := s3repository.NewS3Client(s3Conn)

	jobService := service.NewJobService(worker, redisClient, s3Client)
	jobHandler := handler.NewJobHandler(jobService)
	jobServer := server.NewJobServer(jobHandler)

	if err := jobServer.Run(ctx); err != nil {
		lg.Errorf("job server run error: %v", err)
		return
	}
}
