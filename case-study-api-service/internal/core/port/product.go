//go:generate mockgen -package=mocks -destination=../../../mocks/product_mock.go -source=product.go

package port

import (
	"case-study-api-service/internal/core/models"
	"context"
)

type ProductService interface {
	GetProductByID(ctx context.Context, key string) (*models.GetProductByIDResponse, error)
}

type ProductRepository interface {
	GetProduct(ctx context.Context, key string, parser interface{}) (bool, error)
}
