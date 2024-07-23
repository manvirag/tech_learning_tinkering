package models

type User struct{
  UserId string `json:"user_id"`
  Name string  `json:"name"`
}

type UserRepoInterface interface{
  GetUser(userId string) User
  GetAllUsers() []User
  AddUSer(user User)
}
type UserRepo struct{
  User map[string]User
}

func NewUserRepo() *UserRepo{
  return & UserRepo{
    User: make(map[string]User),
  }
}

func (ur *UserRepo) GetUser(userId string) User {
  return ur.User[userId]
}

func (ur *UserRepo) GetAllUsers() []User {
  users := make([]User, 0, len(ur.User))
  for _, user := range ur.User {
    users = append(users, user)
  }
  return users
}

func (ur *UserRepo) AddUSer(user User) {
  ur.User[user.UserId] = user
}





