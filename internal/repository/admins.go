package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdminsRepo struct {
	db *mongo.Collection
}

func (a AdminsRepo) FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error) {
	result := a.db.FindOne(ctx, bson.M{"email": email, "password": password},
		options.FindOne().SetProjection(bson.M{"password": "0"}))

	var admin domain.Admin
	err := result.Decode(&admin)

	return admin, err
}

func NewAdminsRepo(db *mongo.Database) *AdminsRepo {
	return &AdminsRepo{
		db: db.Collection(adminsCollection),
	}
}
