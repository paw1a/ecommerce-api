package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var products = []domain.Product{
	{
		ID:          primitive.NewObjectID(),
		Name:        "My product",
		Description: "Full product description",
		Price:       1200,
		Categories: []domain.Category{
			{
				ID:          primitive.NewObjectID(),
				Name:        "First category",
				Description: "First category description",
			},
			{
				ID:          primitive.NewObjectID(),
				Name:        "Second category",
				Description: "Second category description",
			},
		},
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "My another product",
		Description: "Full another product description",
		Price:       1500,
		Categories: []domain.Category{
			{
				ID:          primitive.NewObjectID(),
				Name:        "First category",
				Description: "First category description",
			},
			{
				ID:          primitive.NewObjectID(),
				Name:        "Second category",
				Description: "Second category description",
			},
		},
	},
}

func (s *RepositoryTestSuite) TestProductsRepo_FindAll() {
	repo := NewProductsRepo(s.db)
	resultProducts, err := repo.FindAll(context.Background())
	require.NoError(s.T(), err)
	require.Equal(s.T(), products, resultProducts)
}
