package usecases

import (
	"main/models"
	"math"
  "fmt"
)

type ElevatorStratedy interface{
  Run(es *ElevatorSystem ,elevatorId string)
}

type MinWaitingTime struct{
  
}

func (mwt  MinWaitingTime)Run(es *ElevatorSystem, elevatorId string) {
  elevator := es.Elevators[elevatorId]
  for {
     if len(es.RequestQueue) ==0 && len(es.Destinations) == 0 { // stop
           elevator.Direction = 2
           es.Elevators[elevatorId] = elevator
           return ; 
     } else if elevator.Direction == 2 { // no destination, migth have requests
         var minRequest models.Request;
         var nearest int;
         key := ""
         for  reqId , request := range es.RequestQueue {
            diff := int(math.Abs(float64(request.Floor - elevator.Position)))
            if key == "" {
               key = reqId
               nearest = diff
               minRequest = request
            } else if diff < nearest {
              nearest = diff
              minRequest = request
              key = reqId
            }
         }

         if minRequest.Floor < elevator.Position { // go down
              elevator.Direction = 1
         } else if minRequest.Floor > elevator.Position {
              elevator.Direction = 0
         }  else {
           destination := minRequest.User.Notify(elevatorId)
           es.Destinations[destination] = true
           delete(es.RequestQueue, key)
           elevator.Direction = 0
           es.Elevators[elevatorId] = elevator
         }

     } else {
          
         if len(es.RequestQueue) > 0 {
             for reqId, request := range es.RequestQueue {
                 if request.Floor == elevator.Position {
                     destination := request.User.Notify(elevatorId)
                     es.Destinations[destination] = true
                     delete(es.RequestQueue, reqId)
                 }
             }
         }

         if _, ok := es.Destinations[elevator.Position]; ok {
             delete(es.Destinations, elevator.Position)
         }

         if elevator.Direction == 0 { // going up
             if len(es.Destinations) > 0 {
                 for destination := range es.Destinations {
                     if destination > elevator.Position {
                         elevator.Position = elevator.Position+1
                         es.Elevators[elevatorId] = elevator
                         break
                     }
                 }
             } else if len(es.RequestQueue) > 0 {
                 for _, request := range es.RequestQueue {
                     if request.Floor > elevator.Position {
                        elevator.Position = elevator.Position+1
                        es.Elevators[elevatorId] = elevator
                         break
                     }
                 }
             }
         } else if elevator.Direction == 1 { // going down
             if len(es.Destinations) > 0 {
                 for destination := range es.Destinations {
                     if destination < elevator.Position {
                        elevator.Position = elevator.Position-1
                        es.Elevators[elevatorId] = elevator
                        break
                     }
                 }
             } else if len(es.RequestQueue) > 0 {
                 for _, request := range es.RequestQueue {
                     if request.Floor < elevator.Position {
                        elevator.Position = elevator.Position+1
                        es.Elevators[elevatorId] = elevator
                        break
                     }
                 }
             }
         }
     }


    
    fmt.Printf("Elevator ID: %s \n", elevatorId)
    fmt.Printf("Elevator Position: %d\n", elevator.Position)
    switch elevator.Direction {
        case 0:
            fmt.Println("Elevator Direction: ", "Going Up\n\n")
        case 1:
              fmt.Println("Elevator Direction: ", "Going Down\n\n")
        case 2:
              fmt.Println("Elevator Direction: ", "Idle\n\n")
      
    }
    
  }
}