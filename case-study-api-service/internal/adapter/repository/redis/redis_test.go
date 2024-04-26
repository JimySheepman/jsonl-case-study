//go:generate mockgen -package=mocks -destination=../../../../mocks/redis_mock.go -source=redis.go

package repository

import (
	"case-study-api-service/internal/core/models"
	"case-study-api-service/mocks"
	"case-study-api-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localRedisImplMocks struct {
	mockRedisClient *mocks.MockredisClient
}

func _setupRedisImplTest_(t *testing.T) (context.Context, *localRedisImplMocks, *redisImpl) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mock := &localRedisImplMocks{
		mockRedisClient: mocks.NewMockredisClient(ctrl),
	}

	rc := NewRedisClient(mock.mockRedisClient)

	return ctx, mock, rc
}

func TestRedisImpl_GetProduct(t *testing.T) {
	ctx, mock, rc := _setupRedisImplTest_(t)

	tests := []struct {
		name         string
		key          string
		parser       interface{}
		expected     bool
		isError      bool
		expectations func()
	}{
		{
			name:     "redis client get result method return error failure",
			key:      "test-key",
			parser:   &models.GetProductByIDResponse{},
			expected: false,
			isError:  true,
			expectations: func() {
				scmderr := &redis.StringCmd{}
				scmderr.SetErr(customerr.ErrTest)

				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmderr)
			},
		},
		{
			name:     "redis client get result method return redis.nil error failure",
			key:      "test-key",
			parser:   &models.GetProductByIDResponse{},
			expected: false,
			isError:  false,
			expectations: func() {
				scmderr := &redis.StringCmd{}
				scmderr.SetErr(redis.Nil)

				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmderr)
			},
		},
		{
			name:     "json unmarshal error failure",
			key:      "test-key",
			parser:   &models.GetProductByIDResponse{},
			expected: false,
			isError:  true,
			expectations: func() {
				scmd := &redis.StringCmd{}

				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmd)
			},
		},
		{
			name:     "redis client get body method succeed",
			key:      "test-key",
			parser:   &models.GetProductByIDResponse{},
			expected: true,
			isError:  false,
			expectations: func() {
				scmd := &redis.StringCmd{}
				scmd.SetVal(`{"id":234}`)

				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmd)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := rc.GetProduct(ctx, tt.key, tt.parser)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
