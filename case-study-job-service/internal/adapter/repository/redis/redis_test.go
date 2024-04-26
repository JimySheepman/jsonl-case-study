package repository

import (
	"case-study-job-service/internal/core/models"
	"case-study-job-service/mocks"
	"case-study-job-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestRedisImpl_Get(t *testing.T) {
	ctx, mock, rc := _setupRedisImplTest_(t)

	scmd := &redis.StringCmd{}
	scmd.SetVal("test-value")

	tests := []struct {
		name         string
		key          string
		expected     *redis.StringCmd
		expectations func()
	}{
		{
			name:     "redis client get method failure",
			key:      "test-key",
			expected: &redis.StringCmd{},
			expectations: func() {
				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&redis.StringCmd{})
			},
		},
		{
			name:     "redis client get method succeed",
			key:      "test-key",
			expected: scmd,
			expectations: func() {
				val := &redis.StringCmd{}
				val.SetVal("test-value")
				mock.mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(val)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual := rc.Get(ctx, tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRedisImpl_Set(t *testing.T) {
	ctx, mock, rc := _setupRedisImplTest_(t)

	scmd := &redis.StatusCmd{}
	scmd.SetErr(customerr.ErrTest)

	tests := []struct {
		name         string
		key          string
		value        interface{}
		expiration   time.Duration
		expected     *redis.StatusCmd
		expectations func()
	}{
		{
			name:       "redis client set method failure",
			key:        "test-key",
			value:      "test-value",
			expiration: 0,
			expected:   scmd,
			expectations: func() {
				val := &redis.StatusCmd{}
				val.SetErr(customerr.ErrTest)
				mock.mockRedisClient.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(val)
			},
		},
		{
			name:       "redis client set method succeed",
			key:        "test-key",
			value:      "test-value",
			expiration: 0,
			expected:   &redis.StatusCmd{},
			expectations: func() {
				mock.mockRedisClient.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&redis.StatusCmd{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual := rc.Set(ctx, tt.key, tt.value, tt.expiration)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRedisImpl_GetBody(t *testing.T) {
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
			parser:   &models.ProductRequest{},
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
			parser:   &models.ProductRequest{},
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
			parser:   &models.ProductRequest{},
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
			parser:   &models.ProductRequest{},
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

			actual, err := rc.GetBody(ctx, tt.key, tt.parser)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
