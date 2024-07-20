package models

type Elevator struct{
  Id string
  Capacity int
  Position int   
  Direction int // 0 = up, 1 = down, 2 = idle
}

