package dto

import (
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateOrderDTO struct {
	OrderItems  []domain.OrderItem `json:"orderItems" bson:"orderItems"`
	ContactInfo domain.ContactInfo `json:"contactInfo" bson:"contactInfo"`
	UserID      primitive.ObjectID `json:"userID" bson:"userID"`
}

type UpdateOrderDTO struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}

type UpdateOrderInput struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}
