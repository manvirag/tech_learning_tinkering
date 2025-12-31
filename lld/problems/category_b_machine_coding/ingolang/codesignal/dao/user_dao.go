package dao

import . "codesignal/models"

type UserDao struct {
	UserMap      map[string]User
	LoggedInUser map[string]bool
	UserPackage  map[string]Package
}

func NewUserDao() *UserDao {
	return &UserDao{
		UserMap:      make(map[string]User),
		LoggedInUser: make(map[string]bool),
		UserPackage:  make(map[string]Package),
	}
}

// ignoring function right now.
// ignoring func , would be smilary to function
