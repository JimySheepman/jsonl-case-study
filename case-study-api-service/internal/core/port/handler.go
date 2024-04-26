//go:generate mockgen -package=mocks -destination=../../../mocks/handler_mock.go -source=handler.go

package port

import (
	"github.com/labstack/echo/v4"
)

type RestHandlerClient interface {
	HealthCheckHandler(c echo.Context) error
	GetProductByIDHandler(c echo.Context) error
}
