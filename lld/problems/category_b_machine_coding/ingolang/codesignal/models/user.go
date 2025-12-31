package models

type User struct {
	UserId  int
	Name    string
	IsAdmin bool
}

type UserPackage struct {
	UserDetail User
	PackageD   Package
}
