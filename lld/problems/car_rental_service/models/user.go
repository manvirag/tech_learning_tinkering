package models

type UserInterface interface{
  GetId() string
}

type User struct{
  Name string   
  UserId string
}

type Rider struct{
  User
}
func(r Rider)GetId() string{
  return r.UserId
}
type Driver struct{
  User
  Status string
}
func(d Driver)GetId() string{
  return d.UserId
}
func(d Driver)GetStatus() string{
  return d.Status
}
func(d Driver)NotifyForRidy(ride Ride) error{
  return nil
}
func(d Driver)UpdateStatus(driverId string, status string) Driver{
  d.Status = status
  return d
}

type UserRepositoryInterface interface{
  GetUser(userId string) UserInterface
  AddUser(user UserInterface)
}

type UserRepository struct{
  Users map[string]UserInterface
}

func NewUserRepository() *UserRepository{
  return &UserRepository{
    Users: make(map[string]UserInterface),
  }
}

func (r *UserRepository) GetUser(userId string) UserInterface {
  return r.Users[userId]
}

func (r *UserRepository) AddUser(user UserInterface) {
  r.Users[user.GetId()] = user
}