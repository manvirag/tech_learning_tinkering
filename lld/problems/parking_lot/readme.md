- Requirements
  - The parking lot should have multiple levels, each level with a certain number of parking spots.  --> done
  - The parking lot should support different types of vehicles, such as cars, motorcycles, and trucks. --> done
  - Each parking spot should be able to accommodate a specific type of vehicle. --> done
  - The system should assign a parking spot to a vehicle upon entry and release it when the vehicle exits. --> done
  - The system should track the availability of parking spots and provide real-time information to customers. --> done
  - The system should handle multiple entry and exit points and support concurrent access. -> done




- Rough:
  - level -> stratedy type
  - parking spot -> spot -> vechicle
  - parking spot -> level -> list of slots


- models
  - ParkingLotRepo
  - ParkingLotLevel
  - Vehicle -> vechicle interface -> different types.
- usecase
  - ParkingLot