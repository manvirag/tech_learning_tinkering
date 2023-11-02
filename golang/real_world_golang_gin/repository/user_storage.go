package repository

import (
	"context"

	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/domain"
)

type userRepository struct {
	database   string // mongo.Database
	collection string
}

func NewUserRepository(db string, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	// Db client etc code

	user := domain.User{
		ID:       "123",
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "dummy_password",
	}

	return user, nil
}
