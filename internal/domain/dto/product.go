package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
)

type CreateProductDTO struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	Price       float64           `json:"price" binding:"required"`
	Categories  []domain.Category `json:"categories"`
}

type UpdateProductDTO struct {
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Price       *float64          `json:"price"`
	Categories  []domain.Category `json:"categories"`
}

type UpdateProductInput struct {
	Name        string
	Description *string
	Price       *float64
	Categories  []domain.Category
}
