// Code generated by MockGen. DO NOT EDIT.
// Source: redis.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

// MockredisClient is a mock of redisClient interface.
type MockredisClient struct {
	ctrl     *gomock.Controller
	recorder *MockredisClientMockRecorder
}

// MockredisClientMockRecorder is the mock recorder for MockredisClient.
type MockredisClientMockRecorder struct {
	mock *MockredisClient
}

// NewMockredisClient creates a new mock instance.
func NewMockredisClient(ctrl *gomock.Controller) *MockredisClient {
	mock := &MockredisClient{ctrl: ctrl}
	mock.recorder = &MockredisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockredisClient) EXPECT() *MockredisClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockredisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockredisClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockredisClient)(nil).Get), ctx, key)
}
