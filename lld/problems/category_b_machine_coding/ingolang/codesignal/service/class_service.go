package service

import (
	"codesignal/dao"
	"codesignal/models"
	"codesignal/util"
	"errors"
	"fmt"
	"time"
)

type ClassServiceI interface {
}

type ClassService struct {
	classDao *dao.ClassDao
	userDao  *dao.UserDao
	ac       *util.AtomicCounter
}

func NewClassService(cd *dao.ClassDao, ud *dao.UserDao, ac *util.AtomicCounter) *ClassService {
	return &ClassService{classDao: cd, userDao: ud, ac: ac}
}

// ignoring login things in admin routes current
func (cs *ClassService) CreateAdminClassCategory(username string, classType models.ClassCategory) error {
	if _, ok := cs.userDao.UserMap[username]; !ok {
		return errors.New("wrong username")
	}
	usr, _ := cs.userDao.UserMap[username]

	if usr.IsAdmin == false {
		return errors.New("not authorised")
	}
	if _, ok := cs.classDao.AdminClassesTypes[username]; !ok {
		cs.classDao.AdminClassesTypes[username] = make(map[models.ClassCategory]bool)
	}

	cs.classDao.AdminClassesTypes[username][classType] = true
	fmt.Println(cs.classDao.AdminClassesTypes)
	return nil

}
func (cs *ClassService) CreateClasses(username string, classType models.ClassCategory, st time.Time, et time.Time, mxParticipants int) error {
	sch := models.ClassSchedule{
		Id:              cs.ac.GetId(),
		ClassType:       classType,
		StartTime:       st,
		Duration:        int64(et.Sub(st).Hours()),
		MaxParticipants: mxParticipants,
		CreatedBy:       cs.userDao.UserMap[username],
	}
	cs.classDao.ScheduleClasses[sch.Id] = sch
	fmt.Println(cs.classDao.ScheduleClasses)
	return nil
}
