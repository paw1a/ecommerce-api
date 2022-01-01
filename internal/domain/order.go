package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	DeliveredAt time.Time          `json:"deliveredAt" bson:"deliveredAt,omitempty"`
	TotalPrice  float64            `json:"totalPrice" bson:"totalPrice"`
	OrderItems  []OrderItem        `json:"orderItems" bson:"orderItems"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Number    int                `json:"number" bson:"number"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"expiresAt"`
}
