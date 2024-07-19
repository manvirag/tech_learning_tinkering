package models

type Seat interface {
  GetId() string
}

type BasicSeat struct {
	Id string
	Row int
	Column int
  Price int
}



type CheapSeat struct{
  BasicSeat
  Premium int 
}

type PremiumSeat struct{
  BasicSeat
  Premium int
}
func (b BasicSeat) GetId() string {
  return b.Id
}

func (c CheapSeat) GetId() string {
  return c.BasicSeat.Id
}

func (p PremiumSeat) GetId() string {
  return p.BasicSeat.Id
}