package models

type User struct{
  Id string
}

func(u User)Notify(elevatorId string) int{ // lift cam enter the destination
   return 5;
}