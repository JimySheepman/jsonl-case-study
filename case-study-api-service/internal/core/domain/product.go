package domain

import "case-study-api-service/internal/core/models"

type Product struct {
	ID          int64   `db:"Id" json:"id"`
	Title       string  `db:"Title" json:"title"`
	Price       float64 `db:"Price" json:"price"`
	Category    string  `db:"Category" json:"category"`
	Brand       string  `db:"Brand" json:"brand"`
	Url         string  `db:"Url" json:"url"`
	Description string  `db:"Description" json:"description"`
}

func (p Product) ToGetProductByIDResponse() *models.GetProductByIDResponse {
	return &models.GetProductByIDResponse{
		ProductID:   p.ID,
		Title:       p.Title,
		Price:       p.Price,
		Category:    p.Category,
		Brand:       p.Brand,
		Url:         p.Url,
		Description: p.Description,
	}
}
