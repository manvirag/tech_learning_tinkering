Requirements
1. The elevator system should consist of multiple elevators serving multiple floors.
  - x no. of elevators , independent of building storey
2. Each elevator should have a capacity limit and should not exceed it.
  - metadeta for elevator model
    
3. Users should be able to request an elevator from any floor and select a destination floor.
   - use call elevator
   - once come enter destination
   - assuming there is different switch for different elevator
4. The elevator system should efficiently handle user requests and optimize the movement of elevators to minimize waiting time.
   - since there can be multiple request and multiple destination for arrived people at the same time.
7. The system should prioritize requests based on the direction of travel and the proximity of the elevators to the requested floor.
   - giving hint of algorithm to solve the elevator problem.
9. The elevators should be able to handle multiple requests concurrently and process them in an optimal order.
    - go routine use to generate this.
11. The system should ensure thread safety and prevent race conditions when multiple threads interact with the elevators.
    - if require use lock


Rough:

- models
  - elevator
    - weight capacity:
    - id
    - position -> storey
    - direction
  
- usecase
  - ElevatorSystem
    - list of elevators
    - building size
    - elevator working stratedy
    - Request elevator id from floor, multiple concurrent ()
    - destination multiple.()
    - Run()
      - Runing
      - request
      - stop
      - Running
- controller
  - main.go


Algorithm accoriding to requirements.

- at any point of time there will be multiple request for stop and elevator moving in particular direction and have handy destinations.
- if no request and destination -> elevator will stop.
- as soon as we get concurrent request.
  - if elevator not moving or not having desitinations -> choose direction which request is closer.
  - else (or having handy destinations) -> if there is no any destination in direction it is going , it will see the closet request and go in that direction.
    - if have destination -> it will pick all requests if come in that direction until reach the end last in that direction.
    - priority for the already inside persons