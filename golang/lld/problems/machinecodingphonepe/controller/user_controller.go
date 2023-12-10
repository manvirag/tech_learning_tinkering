package controller

import (
	"errors"
	"main/model"
)

type UserController interface {
	CreateUser(userName string, userLoginEmail string, userLoginPassword string) error
	GetUser(userId string) (model.User, error)
}

type UserControllerImpl struct {
	userUsecase model.UserUsecase
}

func NewUserController(userUsecase model.UserUsecase) *UserControllerImpl {
	return &UserControllerImpl{
		userUsecase: userUsecase,
	}
}

func (uc *UserControllerImpl) CreateUser(userName, userLoginEmail, userLoginPassword string) error {

	// Can make this even more complex
	if userName == "" || userLoginEmail == "" || userLoginPassword == "" {
		return errors.New("all fields are required")
	}

	err := uc.userUsecase.CreateUser(userName, userLoginEmail, userLoginPassword)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserControllerImpl) GetUser(userId string) (*model.User, error) {
	if userId == "" {
		return nil, errors.New("invalid user id")
	}

	user, err := uc.userUsecase.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
