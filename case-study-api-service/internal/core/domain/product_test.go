package domain

import (
	"case-study-api-service/internal/core/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToGetProductByIDResponse(t *testing.T) {
	expected := &models.GetProductByIDResponse{
		ProductID:   1,
		Title:       "test",
		Price:       1,
		Category:    "test",
		Brand:       "test",
		Url:         "test",
		Description: "test",
	}

	product := Product{
		ID:          1,
		Title:       "test",
		Price:       1,
		Category:    "test",
		Brand:       "test",
		Url:         "test",
		Description: "test",
	}

	actual := product.ToGetProductByIDResponse()

	assert.Equal(t, expected, actual)
}
