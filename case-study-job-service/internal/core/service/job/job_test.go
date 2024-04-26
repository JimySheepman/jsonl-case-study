package service

import (
	"case-study-job-service/mocks"
	"case-study-job-service/pkg/customerr"
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localJobServiceMocks struct {
	mockWorkerCount       int
	mockProductRepository *mocks.MockProductRepository
	mockRecordRepository  *mocks.MockRecordRepository
}

func _setupJobServiceTest_(t *testing.T) (context.Context, *localJobServiceMocks, *JobService) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mock := &localJobServiceMocks{
		mockWorkerCount:       3,
		mockProductRepository: mocks.NewMockProductRepository(ctrl),
		mockRecordRepository:  mocks.NewMockRecordRepository(ctrl),
	}

	srv := NewJobService(mock.mockWorkerCount, mock.mockProductRepository, mock.mockRecordRepository)

	return ctx, mock, srv
}

func TestJobService_RecordProduct(t *testing.T) {
	ctx, mock, srv := _setupJobServiceTest_(t)

	tests := []struct {
		name         string
		isError      bool
		expectations func()
	}{
		{
			name:    "record repository list objects from bucket failure",
			isError: true,
			expectations: func() {
				mock.mockRecordRepository.EXPECT().ListObjectsFromBucket().Return(nil, customerr.ErrTest)
			},
		},
		{
			name:    "record repository get object failure",
			isError: true,
			expectations: func() {
				mock.mockRecordRepository.EXPECT().ListObjectsFromBucket().Return([]*s3.Object{{}}, nil)
				mock.mockRecordRepository.EXPECT().GetObject(gomock.Any()).Return(nil, customerr.ErrTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.RecordProduct(ctx)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
