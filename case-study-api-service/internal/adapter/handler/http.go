package handler

import (
	"case-study-api-service/internal/core/port"
	"case-study-api-service/pkg/logger"
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

const (
	loggerRequestIdKey  = "requestId"
	appIsHealthyMessage = "program is work"
)

type ProductHandler struct {
	productService port.ProductService
}

func NewProductHandler(recordService port.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: recordService,
	}
}

func (h *ProductHandler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, appIsHealthyMessage)
}

func (h *ProductHandler) GetProductByIDHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	resp, err := h.productService.GetProductByID(ctx, c.Param("id"))
	if err != nil {
		l.Sugar().Errorf("get record by id service error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, resp)
}
