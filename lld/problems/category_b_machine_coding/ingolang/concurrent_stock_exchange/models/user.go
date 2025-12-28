package models

type User struct {
  UserId    string
  UserName  string
	AccountId string
}

type UserRepoInterface interface {
	GetUser(userId string) User
	CreateUser(user User) 
	UpdateUser(user User)
	DeleteUser(userId string)
	GetUserPorfolio(userId string) PortFolio
	GetOrderHis(userId string) []Order
}

type UserRepository struct {
	users       []User
	accountRepo *AccountRepo
}

func NewUserRepository(AccountRepo *AccountRepo) *UserRepository {
	return &UserRepository{
		users:       make([]User, 0),
		accountRepo: AccountRepo,
	}
}

func (ur *UserRepository)CreateUser(user User) {
	ur.users = append(ur.users, user)
}

func (ur *UserRepository)GetUser(userId string) User{
	for _,user := range ur.users {
		if ( user.UserId == userId) {
			return user
			break;
		}
	}
	return User{}
}
