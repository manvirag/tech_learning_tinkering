package dao

import "codesignal/models"

type ClassDao struct {
	// ignore class type in database
	ScheduleClasses   map[int]models.ClassSchedule
	AdminClassesTypes map[string]map[models.ClassCategory]bool
}

func NewClassDao() *ClassDao {
	return &ClassDao{
		ScheduleClasses:   make(map[int]models.ClassSchedule),
		AdminClassesTypes: make(map[string]map[models.ClassCategory]bool),
	}
}

// ignoring func , would be smilary to function
