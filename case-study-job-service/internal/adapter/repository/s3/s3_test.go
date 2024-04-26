package repository

import (
	"case-study-job-service/mocks"
	"case-study-job-service/pkg/customerr"
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type localS3ImplMocks struct {
	mockClient *mocks.Mocks3Client
}

func _setupS3ImplTest_(t *testing.T) (*localS3ImplMocks, *s3Impl) {
	ctrl, _ := gomock.WithContext(context.Background(), t)

	mock := &localS3ImplMocks{
		mockClient: mocks.NewMocks3Client(ctrl),
	}

	s := NewS3Client(mock.mockClient)

	return mock, s
}

func TestS3Impl_GetObject(t *testing.T) {
	mock, s := _setupS3ImplTest_(t)

	tests := []struct {
		name         string
		obj          *s3.Object
		expected     *s3.GetObjectOutput
		isError      bool
		expectations func()
	}{
		{
			name:     "s3 client get object failure",
			obj:      &s3.Object{},
			expected: nil,
			isError:  true,
			expectations: func() {
				mock.mockClient.EXPECT().GetObject(gomock.Any()).Return(nil, customerr.ErrTest)
			},
		},
		{
			name:     "s3 client get object succeed",
			obj:      &s3.Object{},
			expected: &s3.GetObjectOutput{},
			isError:  false,
			expectations: func() {
				mock.mockClient.EXPECT().GetObject(gomock.Any()).Return(&s3.GetObjectOutput{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := s.GetObject(tt.obj)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestS3Impl_ListObjectsFromBucket(t *testing.T) {
	mock, s := _setupS3ImplTest_(t)

	tests := []struct {
		name         string
		expected     []*s3.Object
		isError      bool
		expectations func()
	}{
		{
			name:     "s3 client list objects from bucket failure",
			expected: nil,
			isError:  true,
			expectations: func() {
				mock.mockClient.EXPECT().ListObjectsV2(gomock.Any()).Return(nil, customerr.ErrTest)
			},
		},
		{
			name:     "s3 client list objects from bucket succeed",
			expected: []*s3.Object{{}},
			isError:  false,
			expectations: func() {
				mock.mockClient.EXPECT().ListObjectsV2(gomock.Any()).Return(&s3.ListObjectsV2Output{
					Contents: []*s3.Object{{}},
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := s.ListObjectsFromBucket()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
