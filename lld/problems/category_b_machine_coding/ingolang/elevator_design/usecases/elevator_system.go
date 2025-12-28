package usecases

import (
	"main/models"
	"sync"
)

type ElevatorSystemInterface interface{
  RequesetElevator(floor int, elevatorId string, user models.User)
  
}
type ElevatorSystem struct{
  Elevators map[string]models.Elevator
  BuildingSize int
  RunStratedy ElevatorStratedy
  RequestQueue map[string]models.Request
  Destinations map[int]bool
}

func NewElevatorSystem(buildingSize int, runStratedy ElevatorStratedy) ElevatorSystem{
  return ElevatorSystem{
    BuildingSize: buildingSize,
    Elevators: make(map[string]models.Elevator),
    RunStratedy: runStratedy,
    RequestQueue: make(map[string]models.Request),
    Destinations: make(map[int]bool),
  }
}

func (es *ElevatorSystem)AddElevator(elevator models.Elevator){
  es.Elevators[elevator.Id] = elevator
}

func (es *ElevatorSystem)RequestElevator(reqId string, floor int, elevatorId string, user models.User) {
  es.RequestQueue[reqId] =  models.Request{
    Id: reqId,
    Elevator: es.Elevators[elevatorId],
    User: user,
    Floor: floor,
  }
}

func (es *ElevatorSystem)Run(elevatorId string, wg * sync.WaitGroup) {
   defer wg.Done()
   es.RunStratedy.Run(es, elevatorId)
}

