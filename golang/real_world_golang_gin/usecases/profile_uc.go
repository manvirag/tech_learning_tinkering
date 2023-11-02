package usecases

import (
	"context"

	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
}

func NewProfileUsecase(userRepository domain.UserRepository) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {

	user, err := pu.userRepository.GetByID(c, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Email: user.Email}, nil
}
