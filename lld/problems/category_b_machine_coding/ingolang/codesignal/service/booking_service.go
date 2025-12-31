package service

import (
	"codesignal/dao"
	"codesignal/models"
	"codesignal/util"
	"errors"
	"fmt"
	"time"
)

type BookingService struct {
	bookingDao *dao.BookingDao
	userDao    *dao.UserDao
	classDao   *dao.ClassDao
	ac         *util.AtomicCounter
}

func NewBookingService(bs *dao.BookingDao, ud *dao.UserDao, cd *dao.ClassDao, ac *util.AtomicCounter) *BookingService {
	return &BookingService{bookingDao: bs, userDao: ud, classDao: cd, ac: ac}
}

func (bs *BookingService) BookingClassSchedule(username string, classScheduleId int) error {

	if _, ok := bs.userDao.UserMap[username]; !ok {
		return errors.New("register first")
	} else if upkg, ok := bs.userDao.UserPackage[username]; !ok {
		return errors.New("buy package")
	} else if upkg.GetRemainingClassCount() <= 0 {
		return errors.New("reset your package, no remaining credit")
	} else if sc, ok := bs.classDao.ScheduleClasses[classScheduleId]; !ok {
		return errors.New("wrong schedule class id")
	} else if sc.CurrentParticipants >= sc.MaxParticipants {
		if _, ok := bs.bookingDao.WaitingList[classScheduleId]; ok {
			bs.bookingDao.WaitingList[classScheduleId] = make([]models.User, 0)
		}
		waitingUsers := bs.bookingDao.WaitingList[classScheduleId]
		waitingUsers = append(waitingUsers, bs.userDao.UserMap[username])
		bs.bookingDao.WaitingList[classScheduleId] = waitingUsers
		fmt.Println(bs.bookingDao.WaitingList)
		return errors.New("all classes slot blocked, added you in waiting list")
	}

	if _, ok := bs.userDao.LoggedInUser[username]; !ok {
		return errors.New("login first")
	}

	upkg, _ := bs.userDao.UserPackage[username]
	upkg.SubRemainingClassCount(int(1))
	fmt.Println(upkg)
	bs.userDao.UserPackage[username] = upkg
	sc, _ := bs.classDao.ScheduleClasses[classScheduleId]
	sc.CurrentParticipants = sc.CurrentParticipants + 1
	bs.classDao.ScheduleClasses[classScheduleId] = sc
	bc := models.ClassBooking{
		Id:          bs.ac.GetId(),
		ClassSch:    sc,
		UserDetail:  bs.userDao.UserMap[username],
		BookingTime: time.Now(),
	}
	if _, ok := bs.bookingDao.UserBooking[username]; !ok {
		bs.bookingDao.UserBooking[username] = []models.ClassBooking{}
	}
	existing, _ := bs.bookingDao.UserBooking[username]
	existing = append(existing, bc)
	bs.bookingDao.UserBooking[username] = existing

	fmt.Println(bs.userDao.UserPackage)
	fmt.Println(bs.classDao.ScheduleClasses)
	fmt.Println(bs.bookingDao.UserBooking)
	fmt.Println(bs.bookingDao.WaitingList)
	fmt.Println("............")
	return nil
}

func (bs *BookingService) CancelBooking(bookingId int, username string) error {
	// for show casing just writing without validation, assuming input are  legit
	// logic, validate
	// update state -> wiatinglist, class participant count, user credit

	return nil
}
