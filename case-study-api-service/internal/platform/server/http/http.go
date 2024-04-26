package server

import (
	"case-study-api-service/internal/core/port"
	"case-study-api-service/pkg/config"
	"context"
	"crypto/subtle"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

const pprofEnabled = 1

type RestServer struct {
	e                 *echo.Echo
	restHandlerClient port.RestHandlerClient
}

func NewRestServer(restHandlerClient port.RestHandlerClient) *RestServer {
	return &RestServer{
		e:                 echo.New(),
		restHandlerClient: restHandlerClient,
	}
}

func (s *RestServer) Start() {
	s.setMiddlewares()
	s.setRoutes()

	if err := s.e.Start(config.Cfg.RestServer.Addr); err != nil {
		panic(err)
	}
}

func (s *RestServer) setMiddlewares() {
	s.e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
		middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(config.Cfg.RestServer.Username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(config.Cfg.RestServer.Password)) == 1 {
				return true, nil
			}
			return false, nil
		}),
	)
}

func (s *RestServer) setRoutes() {
	s.e.GET("/healthcheck", s.restHandlerClient.HealthCheckHandler)

	if config.Cfg.RestServer.PprofEnable == pprofEnabled {
		pprof.Register(s.e, "/pprof")
	}

	prefixRouter := s.e.Group("/api/v1")
	prefixRouter.GET("/product/:id", s.restHandlerClient.GetProductByIDHandler)
}

func (s *RestServer) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Fatal(err)
	}
}
