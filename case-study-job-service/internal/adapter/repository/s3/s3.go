//go:generate mockgen -package=mocks -destination=../../../../mocks/s3_mock.go -source=s3.go

package repository

import (
	"case-study-job-service/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Client interface {
	ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error)
}

type s3Impl struct {
	client s3Client
}

func NewS3Client(client s3Client) *s3Impl {
	return &s3Impl{
		client: client,
	}
}

func (i *s3Impl) GetObject(obj *s3.Object) (*s3.GetObjectOutput, error) {
	rawObject, err := i.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.Cfg.AwsS3Config.Bucket),
		Key:    obj.Key,
	})
	if err != nil {
		return nil, err
	}

	return rawObject, nil
}

func (i *s3Impl) ListObjectsFromBucket() ([]*s3.Object, error) {
	resp, err := i.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(config.Cfg.AwsS3Config.Bucket),
	},
	)
	if err != nil {
		return nil, err
	}

	return resp.Contents, nil
}
