package main

import (
	"fmt"
	"main/models"
	"sync"
	"errors"
)

func main() {
	 parkingRepo := models.NewParkingLotRepo()
	 parkingRepo.AddNewLevel(10)
	 parkingRepo.AddNewLevel(10)
	 parkingRepo.AddNewLevel(10)
	 fmt.Printf("%+v \n\n", parkingRepo.SlotCounts)

	 car := models.Car{Number: "0000"}	
   car1 := models.Car{Number: "0001"}	
   bike := models.Bike{Number: "0002"}	
	
	 var wg sync.WaitGroup
	  wg.Add(3)

	 
	 errChan := make(chan error, 3)
	 defer close(errChan)
	
	 go func() {
		 defer wg.Done()
		 level, slot, err := parkingRepo.Park(&car)
		 
		 if(err!=nil){
			 errChan <- errors.New(fmt.Sprintf("Error:%+v,  level: %d, slot: %d , Vechicle: %s", err.Error() ,level, slot, car.Number))
		 } else {
			 fmt.Printf("level: %d, slot: %d , Vechicle: %s\n\n", level, slot, car.Number)
			 errChan <- nil
		 }
	 }()

	
	 go func() {
		 defer wg.Done()
		 level, slot, err := parkingRepo.Park(&car1)
		 if(err!=nil){
			 errChan <- errors.New(fmt.Sprintf("Error:%+v,  level: %d, slot: %d , Vechicle: %s\n\n", err.Error() ,level, slot, car.Number))
		 } else {
			 fmt.Printf("level: %d, slot: %d , Vechicle: %s\n\n", level, slot, car1.Number)
			 errChan <- nil
		 }
	 }()

	
	 go func() {
		 defer wg.Done()
		 level, slot, err := parkingRepo.Park(&bike)
		 if(err!=nil){
			 errChan <- errors.New(fmt.Sprintf("Error:%+v,  level: %d, slot: %d , Vechicle: %s \n\n", err.Error() ,level, slot, car.Number))
		 } else {
			 fmt.Printf("level: %d, slot: %d , Vechicle: %s\n\n", level, slot, bike.Number)
			 errChan <- nil
		 }
	 }()

	
	
	 wg.Wait()
	
	 
	 for i := 0; i < 3; i++ {
		 select {
			 case err := <- errChan:
				 if(err != nil){
					 fmt.Println(err)
				 }
		 }
	 }

	 fmt.Println()
	 fmt.Println(parkingRepo)
	 
	 err := parkingRepo.Release(1,2, &car)
	 if(err != nil ){
		 fmt.Println(err)
	 }


	fmt.Println(parkingRepo)


		 
	
	
}
