package main

import (
	"case-study-api-service/internal/adapter/handler"
	repository "case-study-api-service/internal/adapter/repository/redis"
	service "case-study-api-service/internal/core/service/product"
	server "case-study-api-service/internal/platform/server/http"
	storage "case-study-api-service/internal/platform/storage/redis"
	"case-study-api-service/pkg/config"
	"case-study-api-service/pkg/logger"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	if err := config.LoadConfig("."); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	l := logger.RegisterLogger(config.Cfg.LoggerConfig).Sugar()

	l.Info("case-study-api-service is started...")
	l.Infof("Runtime config %+v", config.Cfg)
	l.Infof("Go runtime is %s", runtime.Version())

	redisConn := storage.RedisConnection()
	redisClient := repository.NewRedisClient(redisConn)
	recordService := service.NewProductService(redisClient)
	recordHandler := handler.NewProductHandler(recordService)

	srv := server.NewRestServer(recordHandler)

	srv.Start()
	l.Info("rest server successfully started")

	signalChan := make(chan os.Signal, 1)
	// syscall package and SIGTERM may be removed for Windows
	// os.Interrupt (SIGINT) and os.Kill (SIGKILL) are only signal values guaranteed to be present on all systems
	// But os.Kill cannot be trapped by a program
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			l.Info("Signal received, shutting down...")
			srv.GracefulShutdown()
			l.Info("rest server graceful shutdown is done")
		}
	}
}
