package usecase

import ("main/models"
        "fmt"
       )

type UberInterface interface{
  RequestRide(source, destination string, userId string) (models.Ride, error)
  GetRide(rideId string) (models.Ride, error)
}

type Uber struct{
  PaymentGateway PaymentGatewayUsecase
  AccountRepo *models.AccountRepo
  UserRepo *models.UserRepository
  FareCalculator FareStrategy
  MatchDriver DriverMatchingStrategy
  RideRepo map[string]models.Ride
}

func NewUber( accountRepo *models.AccountRepo, userRepo *models.UserRepository, fareStratedy FareStrategy, matchDriver DriverMatchingStrategy) *Uber{
  return &Uber{
    PaymentGateway: PaymentGatewayUsecase{},
    AccountRepo: accountRepo,
    UserRepo: userRepo,
    FareCalculator: fareStratedy,
    MatchDriver: matchDriver,
    RideRepo: make(map[string]models.Ride),
  }
}

func (uber *Uber) distance(source, destination string) int32{
  return 50
}

func (uber *Uber) RequestRide(source, destination string, userId string) (models.Ride, error){
  rideCost := uber.FareCalculator.GetFare(source,destination)
  matchedDrivers := uber.MatchDriver.GetMatchingDrivers(source, destination, uber)

  if len(matchedDrivers) == 0{
    return models.Ride{}, fmt.Errorf("No drivers available. Retry in some time")
  }

  ride := models.Ride{
    Id: "ride_123",
    Source: source,
    Destination: destination,
    TotalDistance: uber.distance(source,destination),
    TotalFare: rideCost,
    DriverId: "",
    UserId: userId,
    Status: "Pending",
  }

  // can be optimized further
  
  for _, driverId := range matchedDrivers{
    driver := uber.UserRepo.GetUser(driverId).(models.Driver)
    if driver.GetStatus() == "Available"{
      err := driver.NotifyForRidy(ride)  
      if(err == nil){
        ride.DriverId = driverId
        uber.UserRepo.AddUser(driver.UpdateStatus(driverId, "Busy"));  
        break;
      }
    }
  }

  if ride.DriverId == ""{
    return models.Ride{}, fmt.Errorf("No drivers available. Retry in some time")
  }

  uber.RideRepo[ride.Id] = ride
  return ride , nil
}

func (uber *Uber) GetRide(rideId string) (models.Ride, error) {
	ride, ok := uber.RideRepo[rideId]
	if ok {
		return ride, nil
	}
	return models.Ride{}, fmt.Errorf("Ride not found")
}