package usecases

import (
	"fmt"
	"main/model"
	// "strconv"
	// "time"
)

type UserUsecaseImpl struct {
	userRepository model.UserRepository
}

var userUsecaseImplVar *UserUsecaseImpl

func GetUserUsecase(userRepository model.UserRepository) *UserUsecaseImpl {

	if userUsecaseImplVar == nil {
		userUsecaseImplVar = &UserUsecaseImpl{
			userRepository: userRepository,
		}
	}
	return userUsecaseImplVar
}

func (uc *UserUsecaseImpl) CreateUser(userName, userLoginEmail, UserLoginPassword string) error {

	newUser := model.User{
		ID:       "user123", // for testing
		Name:     userName,
		Email:    userLoginEmail,
		Password: UserLoginPassword,
	}

	fmt.Printf("Sending user to save: %+v", newUser)
	err := uc.userRepository.SaveUser(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecaseImpl) GetUser(userID string) (*model.User, error) {
	user, err := uc.userRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
