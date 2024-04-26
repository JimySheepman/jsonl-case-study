//go:generate mockgen -package=mocks -destination=../../../mocks/service_mock.go -source=service.go

package port

import (
	"context"
)

type ProductService interface {
	RecordProduct(ctx context.Context) error
}
