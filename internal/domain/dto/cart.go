package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
	"time"
)

type CreateCartDTO struct {
	ExpireAt  time.Time         `json:"expireAt"`
	CartItems []domain.CartItem `json:"products"`
}

type UpdateCartDTO struct {
	ExpireAt  *time.Time        `json:"expireAt"`
	CartItems []domain.CartItem `json:"products"`
}

type UpdateCartInput struct {
	ExpireAt  *time.Time        `json:"expireAt"`
	CartItems []domain.CartItem `json:"products"`
}

type UpdateCartItemDTO struct {
	Quantity int64 `json:"quantity"`
}
