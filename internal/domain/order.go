package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderID     string             `json:"orderID" bson:"orderID"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	DeliveredAt time.Time          `json:"deliveredAt" bson:"deliveredAt,omitempty"`
	TotalPrice  float64            `json:"totalPrice" bson:"-"`
	OrderItems  []OrderItem        `json:"orderItems" bson:"orderItems"`
	ContactInfo ContactInfo        `json:"contactInfo" bson:"contactInfo"`
	UserID      primitive.ObjectID `json:"userID" bson:"userID"`
	Status      string             `json:"status" bson:"status"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Quantity  int64              `json:"quantity" bson:"quantity"`
}

type ContactInfo struct {
	Name         string `json:"name" bson:"name"`
	Surname      string `json:"surname" bson:"surname"`
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`
	Address      string `json:"address" bson:"address"`
	OrderComment string `json:"orderComment" bson:"orderComment"`
}
