package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (u UsersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	cursor, err := u.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var usersArray []domain.User
	err = cursor.All(ctx, &usersArray)
	return usersArray, err
}

func (u UsersRepo) FindByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error) {
	result := u.db.FindOne(ctx, bson.M{"_id": userID})

	var user domain.User
	err := result.Decode(&user)

	return user, err
}

func (u UsersRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := u.db.InsertOne(ctx, user)
	return user, err
}

func (u UsersRepo) Update(ctx context.Context, userInput dto.UpdateUserInput, userID primitive.ObjectID) (domain.User, error) {
	updateQuery := bson.M{}

	if userInput.Name != "" {
		updateQuery["name"] = userInput.Name
	}

	_, err := u.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": updateQuery})
	findResult := u.db.FindOne(ctx, bson.M{"_id": userID})

	var user domain.User
	err = findResult.Decode(&user)

	return user, err
}

func (u UsersRepo) Delete(ctx context.Context, userID primitive.ObjectID) error {
	_, err := u.db.DeleteOne(ctx, bson.M{"_id": userID})
	return err
}
