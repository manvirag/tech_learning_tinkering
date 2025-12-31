package models

import "time"

type ClassBooking struct {
	Id          int
	ClassSch    ClassSchedule
	UserDetail  User
	BookingTime time.Time
}
