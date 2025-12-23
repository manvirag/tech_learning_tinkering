package model

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	SaveUser(user User) error
	GetUser(userId string) (*User, error)
}

type UserUsecase interface {
	CreateUser(userName, userLoginEmail, UserLoginPassword string) error
	GetUser(userId string) (*User, error)
}
