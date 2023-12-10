package repository

import (
	"errors"
	"fmt"
	"main/model"
)

// InMemoryUserRepository is a in-memory database implementation of the user
type InMemoryUserRepository struct {
	users map[string]model.User
}

var inMemoryUserRepository *InMemoryUserRepository

func NewInMemoryUserRepository() *InMemoryUserRepository {
	if inMemoryUserRepository == nil {
		inMemoryUserRepository = &InMemoryUserRepository{
			users: make(map[string]model.User),
		}
	}
	return inMemoryUserRepository
}

func (r *InMemoryUserRepository) SaveUser(user model.User) error {
	if _, exists := r.users[user.ID]; exists {
		fmt.Println("This user already exists")
		return nil
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetUser(userId string) (*model.User, error) {
	user, exists := r.users[userId]
	if !exists {
		return nil, errors.New("user doesn not found. Please create user first")
	}
	return &user, nil
}
