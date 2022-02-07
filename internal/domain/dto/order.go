package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateOrderDTO struct {
	OrderItems  []domain.OrderItem `json:"-"`
	ContactInfo domain.ContactInfo `json:"contactInfo"`
	UserID      primitive.ObjectID `json:"-"`
}

type UpdateOrderDTO struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}

type UpdateOrderInput struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}
