package storage

import (
	"case-study-job-service/pkg/config"
	"case-study-job-service/pkg/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"sync"
)

var (
	client *s3.S3
	once   sync.Once
)

func S3Connection() *s3.S3 {
	l := logger.GetLogger().Sugar()

	once.Do(func() {
		l.Info("trying to create session s3 server")
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(config.Cfg.AwsS3Config.Region),
			Credentials: credentials.NewStaticCredentials(
				config.Cfg.AwsS3Config.SecretId,
				config.Cfg.AwsS3Config.SecretKey,
				config.Cfg.AwsS3Config.Token),
		})
		if err != nil {
			panic(err)
		}
		l.Info("create session s3 successful")

		client = s3.New(sess)
	})

	return client
}
