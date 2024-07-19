package main

import (
	"fmt"
	"main/models"
	"main/usecases"
)

func main() {
     theatreOne := models.Theatre{
			 Name: "Wakad",
			 Id: "Pune_1",
			 Seats: []models.Seat{
				 models.BasicSeat{
					 Id: "11",
					 Row: 1,
					 Column: 1,
					 Price: 10,
				 },
				 models.BasicSeat{
						Id: "12",
						Row: 1,
						Column: 2,
						Price: 10,
					},
				 models.CheapSeat{
					 BasicSeat: models.BasicSeat {
							 Id: "21",
								 Row: 2,
								 Column: 1,
								 Price: 10,
					 },
					 Premium: 5,
				 },
				 models.CheapSeat{
						BasicSeat: models.BasicSeat{
								Id: "31",
									Row: 3,
									Column: 1,
									Price: 10,
						},
						Premium: 7,
				 },
			 },
			 Slots: []models.Slot{
				 models.Slot{
					 Id: "1",
					 StartTime: 1,
					 EndTime: 4,
				 },
				 models.Slot{
					 Id: "2",
					 StartTime : 7,
					 EndTime: 10,
				 },
			 },
		 }

		 inventory := models.NewInventory()
	   
	   fmt.Println(inventory.GetTheatre(theatreOne.Id))

	   dhoom := models.Movie{
			  Name: "DHOOM",
			  Id: "Dhoom",
		 }

	   theatreOne.Slots[0].MovieDetail = dhoom

		 err := inventory.AddTheatre(theatreOne)
			if (err != nil){
				fmt.Println(err)
		}

	  // fmt.Println(inventory.GetTheatre(theatreOne.Id))
		fmt.Println(inventory)
		
		mmt :=usecases.NewMMT(inventory)
	  mmt.BookMovie("Pune_1", "1", "11")

	  fmt.Println(inventory)
		
}
