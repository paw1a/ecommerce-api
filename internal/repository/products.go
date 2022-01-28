package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepo struct {
	db *mongo.Collection
}

func (p ProductsRepo) FindAll(ctx context.Context) ([]domain.Product, error) {
	cursor, err := p.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var productArray []domain.Product
	err = cursor.All(ctx, &productArray)
	return productArray, err
}

func (p ProductsRepo) FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error) {
	result := p.db.FindOne(ctx, bson.M{"_id": productID})

	var product domain.Product
	err := result.Decode(&product)

	return product, err
}

func (p ProductsRepo) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.db.InsertOne(ctx, product)
	return product, err
}

func (p ProductsRepo) Update(ctx context.Context, productInput dto.UpdateProductInput, productID primitive.ObjectID) (domain.Product, error) {
	updateQuery := bson.M{}

	if productInput.Name != "" {
		updateQuery["name"] = productInput.Name
	}

	if productInput.Description != nil {
		updateQuery["description"] = productInput.Description
	}

	if productInput.Price != nil {
		updateQuery["price"] = productInput.Price
	}

	if productInput.Categories != nil {
		updateQuery["categories"] = productInput.Categories
	}

	_, err := p.db.UpdateOne(ctx, bson.M{"_id": productID}, bson.M{"$set": updateQuery})
	findResult := p.db.FindOne(ctx, bson.M{"_id": productID})

	var product domain.Product
	err = findResult.Decode(&product)

	return product, err
}

func (p ProductsRepo) Delete(ctx context.Context, productID primitive.ObjectID) error {
	_, err := p.db.DeleteOne(ctx, bson.M{"_id": productID})
	return err
}

func NewProductsRepo(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{
		db: db.Collection(productsCollection),
	}
}
