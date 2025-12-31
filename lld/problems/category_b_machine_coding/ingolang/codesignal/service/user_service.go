package service

import (
	"codesignal/dao"
	"codesignal/models"
	"codesignal/util"
	"errors"
	"fmt"
)

// assume name -> username -> unique
// not handling company manager and all asusming internal api
type UserServiceI interface {
	RegisterUser(username string) error
	AddAdminUser(username string) error // no access to user
	LogInUser(username string) error
	LogOutUser(username string) error
	BuyPackage(username string, packageType models.PackageCategory) error
}

type UserService struct {
	userDao *dao.UserDao
	ac      *util.AtomicCounter
}

func NewUserService(ud *dao.UserDao, ac *util.AtomicCounter) *UserService {
	return &UserService{
		userDao: ud,
		ac:      ac,
	}
}

func (us *UserService) RegisterUser(username string) error {
	if _, ok := us.userDao.UserMap[username]; ok {
		return errors.New("already exist username")
	}
	usr := models.User{
		UserId:  us.ac.GetId(),
		Name:    username,
		IsAdmin: false,
	}

	us.userDao.UserMap[username] = usr
	fmt.Println(us.userDao.UserMap)
	return nil
}

func (us *UserService) AddAdminUser(username string) error {
	if _, ok := us.userDao.UserMap[username]; ok {
		return errors.New("already exist username")
	}
	usr := models.User{
		UserId:  us.ac.GetId(),
		Name:    username,
		IsAdmin: true,
	}
	us.userDao.UserMap[username] = usr
	fmt.Println(us.userDao.UserMap)
	return nil
}
func (us *UserService) LogInUser(username string) error {
	if _, ok := us.userDao.UserMap[username]; !ok {
		return errors.New("register first")
	}
	us.userDao.LoggedInUser[username] = true
	fmt.Println(us.userDao.LoggedInUser)
	return nil
}
func (us *UserService) LogOutUser(username string) error {
	if _, ok := us.userDao.UserMap[username]; !ok {
		return errors.New("register first")
	}
	delete(us.userDao.LoggedInUser, username)
	fmt.Println(us.userDao.LoggedInUser)
	return nil
}

// assuming no override validation
func (us *UserService) BuyPackage(username string, packageType models.PackageCategory) error {
	pkg, err := models.PackageFactory(packageType)
	if err != nil {
		return errors.New("wrong package type")
	}

	us.userDao.UserPackage[username] = pkg
	fmt.Println(us.userDao.UserPackage)

	return nil
}
