package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
)

type CreateCartDTO struct {
	Products []domain.CartItem `json:"products" bson:"products"`
}

type UpdateCartDTO struct {
	Products []domain.CartItem `json:"products"`
}

type UpdateCartInput struct {
	Products []domain.CartItem `json:"products"`
}
