package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartsRepo struct {
	db *mongo.Collection
}

func (c *CartsRepo) FindAll(ctx context.Context) ([]domain.Cart, error) {
	cursor, err := c.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var cartArray []domain.Cart
	err = cursor.All(ctx, &cartArray)
	return cartArray, err
}

func (c *CartsRepo) FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error) {
	result := c.db.FindOne(ctx, bson.M{"_id": cartID})

	var cart domain.Cart
	err := result.Decode(&cart)

	return cart, err
}

func (c *CartsRepo) Create(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	cart.ID = primitive.NewObjectID()
	_, err := c.db.InsertOne(ctx, cart)
	return cart, err
}

func (c *CartsRepo) Update(ctx context.Context, cartInput dto.UpdateCartInput, cartID primitive.ObjectID) (domain.Cart, error) {
	updateQuery := bson.M{}

	if cartInput.Products != nil {
		updateQuery["products"] = cartInput.Products
	}

	if cartInput.ExpireAt != nil {
		updateQuery["expireAt"] = cartInput.ExpireAt
	}

	_, err := c.db.UpdateOne(ctx, bson.M{"_id": cartID}, bson.M{"$set": updateQuery})
	findResult := c.db.FindOne(ctx, bson.M{"_id": cartID})

	var cart domain.Cart
	err = findResult.Decode(&cart)

	return cart, err
}

func (c *CartsRepo) Delete(ctx context.Context, cartID primitive.ObjectID) error {
	_, err := c.db.DeleteOne(ctx, bson.M{"_id": cartID})
	return err
}

func NewCartsRepo(db *mongo.Database) *CartsRepo {
	collection := db.Collection(cartsCollection)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"expireAt": 1},
		Options: options.Index().SetExpireAfterSeconds(0),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("unable to create cart collection index, %v", err)
	}

	return &CartsRepo{
		db: collection,
	}
}
