package models

type Request struct{
  Id string
  Elevator Elevator
  User User
  Floor int
}