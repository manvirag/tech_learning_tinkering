package usecase

import (
	"fmt"
	"main/models"
)

type UberInterface interface {
	RequestRide(source, destination string, userId string) (models.Ride, error)
	GetRide(rideId string) (string, error)
	AcceptRide(rideId string, driverId string) (string, error)
	StartRidy(rideId string) (string, error)
	CompleteRide(rideId string) (string, error)
}

type Uber struct {
	PaymentGateway PaymentGatewayUsecase
	AccountRepo    *models.AccountRepo
	UserRepo       *models.UserRepository
	FareCalculator FareStrategy
	MatchDriver    DriverMatchingStrategy
	RideRepo       map[string]models.Ride
}

func NewUber(accountRepo *models.AccountRepo, userRepo *models.UserRepository, fareStratedy FareStrategy, matchDriver DriverMatchingStrategy) *Uber {
	return &Uber{
		PaymentGateway: PaymentGatewayUsecase{},
		AccountRepo:    accountRepo,
		UserRepo:       userRepo,
		FareCalculator: fareStratedy,
		MatchDriver:    matchDriver,
		RideRepo:       make(map[string]models.Ride),
	}
}

func (uber *Uber) distance(source, destination string) int32 {
	return 50
}

func (uber *Uber) RequestRide(source, destination string, userId string) (string, error) {
	rideCost := uber.FareCalculator.GetFare(source, destination)
	matchedDrivers := uber.MatchDriver.GetMatchingDrivers(source, destination, uber)

	if len(matchedDrivers) == 0 {
		return "", fmt.Errorf("No drivers available. Retry in some time")
	}

	ride := models.Ride{
		Id:            "ride_123",
		Source:        source,
		Destination:   destination,
		TotalDistance: uber.distance(source, destination),
		TotalFare:     rideCost,
		DriverId:      "",
		UserId:        userId,
		Status:        "Pending",
	}

	// can be optimized further

	for _, driverId := range matchedDrivers {
		driver := uber.UserRepo.GetUser(driverId).(models.Driver)
		err := driver.NotifyForRidy(ride)
		if err != nil {
			fmt.Println(err)
		}

	}

	uber.RideRepo[ride.Id] = ride
	return ride.Id, nil
}

func (uber *Uber) GetRide(rideId string) (models.Ride, error) {
	ride, ok := uber.RideRepo[rideId]
	if ok {
		return ride, nil
	}
	return models.Ride{}, fmt.Errorf("Ride not found")
}

func (uber *Uber) AcceptRide(rideId string, driverId string) (string, error) {
	ride, err := uber.GetRide(rideId)
	if err != nil {
		return "", err
	}

	if ride.Status != "Pending" {
		return "", fmt.Errorf("Ride is not pending")
	}

	ride.DriverId = driverId
	ride.Status = "Accepted"

	uber.RideRepo[rideId] = ride
	return "Ride accepted", nil
}

func (uber *Uber) StartRidy(rideId string) (string, error) {
	ride, err := uber.GetRide(rideId)
	if err != nil {
		return "", err
	}

	if ride.Status != "Accepted" {
		return "", fmt.Errorf("Ride is not accepted")
	}

	ride.Status = "Started"
	uber.RideRepo[rideId] = ride
	return "Ride started", nil
}

func (uber *Uber) CompleteRide(rideId string) (string, error) {
	ride, err := uber.GetRide(rideId)
	if err != nil {
		return "", err
	}

	if ride.Status != "Started" {
		return "", fmt.Errorf("Ride is not started")
	}

	ride.Status = "Completed"
	uber.RideRepo[rideId] = ride
	return "Ride completed", nil
}
