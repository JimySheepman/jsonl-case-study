package handler

import (
	mock_port "case-study-job-service/mocks"
	"case-study-job-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localJobHandlerMocks struct {
	mockProductService *mock_port.MockProductService
}

func _setupJobHandlerTest_(t *testing.T) (context.Context, *localJobHandlerMocks, *JobHandler) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &localJobHandlerMocks{
		mockProductService: mock_port.NewMockProductService(ctrl),
	}

	hdlr := NewJobHandler(mocks.mockProductService)

	return ctx, mocks, hdlr

}

func TestJobHandler_Run(t *testing.T) {
	ctx, mocks, hdlr := _setupJobHandlerTest_(t)

	tests := []struct {
		name         string
		isError      bool
		expectations func()
	}{
		{
			name:    "product service record product failure",
			isError: true,
			expectations: func() {
				mocks.mockProductService.EXPECT().RecordProduct(gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:    "job handler run method succeed",
			isError: false,
			expectations: func() {
				mocks.mockProductService.EXPECT().RecordProduct(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := hdlr.Run(ctx)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
