package models

type VehicleInterface interface{
  GetVehicleNumber() string
}
type Car struct {
   Number string 
}

func (c *Car) GetVehicleNumber() string{
  return c.Number
}

type Bike struct {
   Number string 
}

func (c *Bike) GetVehicleNumber() string{
  return c.Number
}



