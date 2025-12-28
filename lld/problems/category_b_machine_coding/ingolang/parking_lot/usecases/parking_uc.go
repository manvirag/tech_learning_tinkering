package usecases

import "main/models"

type ParkingLotInterface interface{
  
}

type ParkingLot struct{
  parkingRepo *models.ParkingLotRepo
}