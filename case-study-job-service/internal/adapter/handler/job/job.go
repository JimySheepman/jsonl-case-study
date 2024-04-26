package handler

import (
	"case-study-job-service/internal/core/port"
	"context"
)

type JobHandler struct {
	productService port.ProductService
}

func NewJobHandler(productService port.ProductService) *JobHandler {
	return &JobHandler{
		productService: productService,
	}
}

func (h JobHandler) Run(ctx context.Context) error {
	return h.productService.RecordProduct(ctx)
}
