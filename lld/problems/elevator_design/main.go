package main

import (
	"fmt"
	"main/models"
	"main/usecases"
	"sync"
)

func main() {
	var wg sync.WaitGroup;
	es := usecases.NewElevatorSystem(10,usecases.MinWaitingTime{},)
	fmt.Println(es)
	
	es.AddElevator(models.Elevator{
		Id: "1",
		Capacity: 100,
		Position: 0,
		Direction: 2,
	})

	fmt.Println(es)

	
	wg.Add(1)
	go es.Run("1", &wg)

  
	es.RequestElevator("Req_1", 1, "1", models.User{
		Id: "user_1",
	})

	es.RequestElevator("Req_2", 3, "1", models.User{
		Id: "user_2",
	})


	es.RequestElevator("Req_3", 0, "1", models.User{
		Id: "user_3",
	})

	fmt.Println(es.RequestQueue)

	
	
	wg.Wait()

	
}
