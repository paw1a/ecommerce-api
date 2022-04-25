package repository

import (
	"context"
	"github.com/paw1a/ecommerce-api/internal/config"
	"github.com/paw1a/ecommerce-api/pkg/database/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite

	db     *mongo.Database
	config *config.Config
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupSuite() {
	s.config = config.GetConfig("../../config/config.yml")

	client, err := mongodb.NewClient(context.Background(),
		s.config.Test.DBURI, s.config.Test.DBUsername, s.config.Test.DBPassword)
	if err != nil {
		s.FailNow("Failed to connect to mongo", err)
	}

	s.db = client.Database(s.config.Test.Database)

	err = s.initDB()
	if err != nil {
		s.FailNow("can't create products collection for tests")
	}
}

func (s *RepositoryTestSuite) initDB() error {
	for _, product := range products {
		_, err := s.db.Collection(productsCollection).InsertOne(context.Background(), product)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RepositoryTestSuite) TearDownSuite() {
	err := s.db.Drop(context.Background())
	if err != nil {
		s.FailNow("can't drop database with name: ", s.db.Name())
	}

	err = s.db.Client().Disconnect(context.Background())
	if err != nil {
		s.FailNow("can't disconnect from db")
	}
}
