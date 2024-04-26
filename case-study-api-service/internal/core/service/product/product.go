package service

import (
	"case-study-api-service/internal/core/domain"
	"case-study-api-service/internal/core/models"
	"case-study-api-service/internal/core/port"
	"context"
	"fmt"
)

type ProductService struct {
	productRepository port.ProductRepository
}

func NewProductService(productRepository port.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s ProductService) GetProductByID(ctx context.Context, key string) (*models.GetProductByIDResponse, error) {
	resp := &domain.Product{}
	ok, err := s.productRepository.GetProduct(ctx, key, resp)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("product is not found")
	}

	return resp.ToGetProductByIDResponse(), nil
}
