package usecases

import (
	"errors"
	"main/models"
)

type MMTInterface interface{
  BookMovie(theatreId string, slotId string, seatId string) error
}

type MMT struct{
  inventory *models.Inventory
}

func NewMMT(inventor *models.Inventory) *MMT{
  return & MMT{
    inventory: inventor,
  }
}

func (mmt *MMT)BookMovie(theatreId string, slotId string, seatId string) error {
  theatre, err := mmt.inventory.GetTheatre(theatreId)
  if err != nil {
    return err
  }
  slotExist := false
  for _,slot := range theatre.Slots{
     if( slot.Id == slotId){
          slotExist = true
          break
     }
  }
  if(slotExist == false){
    return errors.New("Wrong slot")
  }
  
  for _, seat := range mmt.inventory.AvailableSeats[theatre.Id][slotId]{
      if(seat.GetId() == seatId){
          err = mmt.inventory.DeleteSeat(theatre.Id, slotId, seatId)
          if(err != nil ){
            return err
          }
          return nil
      }
  }
  
  return errors.New("Seat does not exist")
}
  


