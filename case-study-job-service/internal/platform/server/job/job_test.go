package server

import (
	mock_port "case-study-job-service/mocks"
	"case-study-job-service/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localJobServerMocks struct {
	mockJobHandlerClient *mock_port.MockJobHandlerClient
}

func _setupJobServerTest_(t *testing.T) (context.Context, *localJobServerMocks, *JobServer) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &localJobServerMocks{
		mockJobHandlerClient: mock_port.NewMockJobHandlerClient(ctrl),
	}

	srv := NewJobServer(mocks.mockJobHandlerClient)

	return ctx, mocks, srv

}

func TestJobServer_Run(t *testing.T) {
	ctx, mocks, srv := _setupJobServerTest_(t)

	tests := []struct {
		name         string
		isError      bool
		expectations func()
	}{
		{
			name:    "product service record product failure",
			isError: true,
			expectations: func() {
				mocks.mockJobHandlerClient.EXPECT().Run(gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:    "job server run method succeed",
			isError: false,
			expectations: func() {
				mocks.mockJobHandlerClient.EXPECT().Run(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.Run(ctx)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
