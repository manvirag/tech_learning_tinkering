package models

import ("errors"
        "sync"
       )


type InventoryInterface interface{
  GetTheatre(id string) (Theatre, error) 
  AddTheatre(theatre Theatre) error
  DeleteSeat(theatreId string, slotId string, seatId string) error
} 

type Inventory struct{     // MMT DB
   Theatres map[string]Theatre
   AvailableSeats map[string]map[string][]Seat
   sync.RWMutex
}

func NewInventory() *Inventory{
  return & Inventory{
    Theatres: make(map[string]Theatre),
    AvailableSeats: map[string]map[string][]Seat{},
  }
}

func (i *Inventory)GetTheatre(id string) (Theatre, error) {
    if theatre, ok := i.Theatres[id]; ok{
      return theatre , nil
    }

    return Theatre{}, errors.New("Wrong theatre Id")
}

func (i *Inventory) AddTheatre(theatre Theatre) error {
	if _, ok := i.Theatres[theatre.Id]; ok {
		return errors.New("Theatre with this ID already exists")
	}
	i.Theatres[theatre.Id] = theatre
  i.AvailableSeats[theatre.Id] = map[string][]Seat{}
  for _, slot := range theatre.Slots {
      i.AvailableSeats[theatre.Id][slot.Id] = theatre.Seats
  }
	return nil
}

func (i *Inventory) DeleteSeat(theatreId string, slotId string, seatId string) error {
	i.Lock()
	defer i.Unlock()

	for idx, seat := range i.AvailableSeats[theatreId][slotId] {
		if seat.GetId() == seatId {
			i.AvailableSeats[theatreId][slotId] = append(i.AvailableSeats[theatreId][slotId][:idx], i.AvailableSeats[theatreId][slotId][idx+1:]...)
			return nil
		}
	}
	return errors.New("Seat with this ID does not exist in this Slot")
}


