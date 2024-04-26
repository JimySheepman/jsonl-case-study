package service

import (
	"case-study-api-service/internal/core/models"
	"case-study-api-service/mocks"
	"case-study-api-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localProductServiceMocks struct {
	mockProductRepository *mocks.MockProductRepository
}

func _setupProductServiceTest_(t *testing.T) (context.Context, *localProductServiceMocks, *ProductService) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mock := &localProductServiceMocks{
		mockProductRepository: mocks.NewMockProductRepository(ctrl),
	}

	srv := NewProductService(mock.mockProductRepository)

	return ctx, mock, srv
}

func TestProductService_GetProductByID(t *testing.T) {
	ctx, mock, srv := _setupProductServiceTest_(t)

	tests := []struct {
		name         string
		key          string
		expected     *models.GetProductByIDResponse
		isError      bool
		expectations func()
	}{
		{
			name:     "GetProduct failure",
			key:      "test-key",
			expected: nil,
			isError:  true,
			expectations: func() {
				mock.mockProductRepository.EXPECT().GetProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, customerr.ErrTest)
			},
		},
		{
			name:     "GetProduct response is false failure",
			key:      "test-key",
			expected: nil,
			isError:  true,
			expectations: func() {
				mock.mockProductRepository.EXPECT().GetProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
			},
		},
		{
			name:     "GetProductByID succeed",
			key:      "test-key",
			expected: &models.GetProductByIDResponse{},
			isError:  false,
			expectations: func() {
				mock.mockProductRepository.EXPECT().GetProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.GetProductByID(ctx, tt.key)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
