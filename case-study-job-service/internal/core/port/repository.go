//go:generate mockgen -package=mocks -destination=../../../mocks/repository_mock.go -source=repository.go

package port

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/redis/go-redis/v9"
	"time"
)

type ProductRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	GetBody(ctx context.Context, key string, parser interface{}) (bool, error)
}

type RecordRepository interface {
	GetObject(obj *s3.Object) (*s3.GetObjectOutput, error)
	ListObjectsFromBucket() ([]*s3.Object, error)
}
