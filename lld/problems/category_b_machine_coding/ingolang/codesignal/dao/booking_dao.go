package dao

import "codesignal/models"

type BookingDao struct {
	UserBooking map[string][]models.ClassBooking
	WaitingList map[int][]models.User // schedule id vs list of users in wiating.
}

func NewBookingDao() *BookingDao {
	return &BookingDao{
		UserBooking: make(map[string][]models.ClassBooking),
		WaitingList: make(map[int][]models.User),
	}
}
