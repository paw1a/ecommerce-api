package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProductDTO struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	Price       float64           `json:"price" binding:"required"`
	Categories  []domain.Category `json:"categories"`
}

type UpdateProductDTO struct {
	ID          primitive.ObjectID `json:"-"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	Price       *float64           `json:"price"`
	TotalRating *float32           `json:"totalRating,omitempty"`
	Categories  []domain.Category  `json:"categories"`
}

type UpdateProductInput struct {
	ID          primitive.ObjectID
	Name        string
	Description *string
	Price       *float64
	TotalRating *float32
	Categories  []domain.Category
}
