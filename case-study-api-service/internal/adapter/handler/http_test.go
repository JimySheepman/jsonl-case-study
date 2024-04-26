package handler

import (
	"bytes"
	"case-study-api-service/internal/core/models"
	"case-study-api-service/mocks"
	"case-study-api-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type localProductHandlerMocks struct {
	mockProductService *mocks.MockProductService
}

func _setupProductHandlerTest_(t *testing.T) (*echo.Echo, *localProductHandlerMocks, *ProductHandler) {
	ctrl, _ := gomock.WithContext(context.Background(), t)

	mock := &localProductHandlerMocks{
		mockProductService: mocks.NewMockProductService(ctrl),
	}

	hdlr := NewProductHandler(mock.mockProductService)
	e := echo.New()

	return e, mock, hdlr
}

func TestProductHandler_HealthCheckHandler(t *testing.T) {
	e, _, hdlr := _setupProductHandlerTest_(t)

	tests := []struct {
		name string
	}{
		{
			name: "HealthCheckHandler succeed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := hdlr.HealthCheckHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestProductHandler_GetProductByIDHandler(t *testing.T) {
	e, mock, hdlr := _setupProductHandlerTest_(t)

	tests := []struct {
		name         string
		body         []byte
		expectations func()
	}{
		{
			name: "GetProductByID failure",
			body: nil,
			expectations: func() {
				mock.mockProductService.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, customerr.ErrTest)
			},
		},
		{
			name: "GetProductByIDHandler succeed",
			body: nil,
			expectations: func() {
				mock.mockProductService.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.GetProductByIDResponse{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := hdlr.GetProductByIDHandler(ctx)
			assert.Nil(t, err)
		})
	}
}
