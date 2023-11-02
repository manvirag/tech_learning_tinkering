package domain

import (
	"context"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepository interface {
	GetByID(c context.Context, id string) (User, error)
}
