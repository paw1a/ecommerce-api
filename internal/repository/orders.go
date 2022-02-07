package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrdersRepo struct {
	db *mongo.Collection
}

func (p *OrdersRepo) FindAll(ctx context.Context) ([]domain.Order, error) {
	cursor, err := p.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var orderArray []domain.Order
	err = cursor.All(ctx, &orderArray)
	return orderArray, err
}

func (p *OrdersRepo) FindByID(ctx context.Context, orderID primitive.ObjectID) (domain.Order, error) {
	result := p.db.FindOne(ctx, bson.M{"_id": orderID})

	var order domain.Order
	err := result.Decode(&order)

	return order, err
}

func (p *OrdersRepo) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Order, error) {
	cursor, err := p.db.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}

	var orderArray []domain.Order
	err = cursor.All(ctx, &orderArray)
	return orderArray, err
}

func (p *OrdersRepo) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	order.ID = primitive.NewObjectID()
	_, err := p.db.InsertOne(ctx, order)
	return order, err
}

func (p *OrdersRepo) Update(ctx context.Context, productInput dto.UpdateOrderInput, orderID primitive.ObjectID) (domain.Order, error) {
	updateQuery := bson.M{}

	if productInput.Status != "" {
		updateQuery["status"] = productInput.Status
	}

	if !productInput.DeliveredAt.IsZero() {
		updateQuery["deliveredAt"] = productInput.DeliveredAt
	}

	_, err := p.db.UpdateOne(ctx, bson.M{"_id": orderID}, bson.M{"$set": updateQuery})
	findResult := p.db.FindOne(ctx, bson.M{"_id": orderID})

	var order domain.Order
	err = findResult.Decode(&order)

	return order, err
}

func (p *OrdersRepo) Delete(ctx context.Context, orderID primitive.ObjectID) error {
	_, err := p.db.DeleteOne(ctx, bson.M{"_id": orderID})
	return err
}

func NewOrdersRepo(db *mongo.Database) *OrdersRepo {
	return &OrdersRepo{
		db: db.Collection(ordersCollection),
	}
}
