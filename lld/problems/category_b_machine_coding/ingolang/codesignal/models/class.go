package models

import "time"

type ClassCategory string

const (
	Gym  ClassCategory = "GYM"
	Yoga ClassCategory = "YOGA"
)

type Class interface {
	GetType() ClassCategory
}

type GymClass struct {
	Type ClassCategory
}

func (pp GymClass) GetType() ClassCategory {
	return pp.Type
}

type YogaClass struct {
	Type ClassCategory
}

func (pp YogaClass) GetType() ClassCategory {
	return pp.Type
}

type ClassSchedule struct {
	Id                  int
	ClassType           ClassCategory
	StartTime           time.Time
	Duration            int64 // hours
	MaxParticipants     int
	CurrentParticipants int
	CreatedBy           User
}

func ClassFactory(category ClassCategory) (Class, error) { // current hardcoding class count
	switch category {
	case Gym:
		return GymClass{
			Type: Gym,
		}, nil
	case Yoga:
		return YogaClass{
			Type: Yoga,
		}, nil
	}

	return nil, nil
}
