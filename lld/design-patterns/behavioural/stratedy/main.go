package main

import "fmt"

// DriveStrategy defines the interface for various driving strategies.
type DriveStrategy interface {
	Drive()
}

// NormalDriverStrategy is a concrete implementation of DriveStrategy for normal driving.
type NormalDriverStrategy struct{}

func (s *NormalDriverStrategy) Drive() {
	fmt.Println("Normal drive strategy")
}

// SportDriveStrategy is a concrete implementation of DriveStrategy for sports driving.
type SportDriveStrategy struct{}

func (s *SportDriveStrategy) Drive() {
	fmt.Println("Sports drive strategy")
}

// Vehicle is a class that uses a DriveStrategy to perform driving.
type Vehicle struct {
	driveStrategy DriveStrategy
}

func (v *Vehicle) Drive() {
	v.driveStrategy.Drive()
}

// SportsVehicle is a specific type of vehicle with a sports driving strategy.
type SportsVehicle struct {
	Vehicle
}

func NewSportsVehicle() *SportsVehicle {
	return &SportsVehicle{Vehicle: Vehicle{driveStrategy: &SportDriveStrategy{}}}
}

// OffRoadVehicle is a specific type of vehicle with a normal driving strategy.
type OffRoadVehicle struct {
	Vehicle
}

func NewOffRoadVehicle() *OffRoadVehicle {
	return &OffRoadVehicle{Vehicle: Vehicle{driveStrategy: &NormalDriverStrategy{}}}
}

// GoodsVehicle is a specific type of vehicle with a normal driving strategy.
type GoodsVehicle struct {
	Vehicle
}

func NewGoodsVehicle() *GoodsVehicle {
	return &GoodsVehicle{Vehicle: Vehicle{driveStrategy: &NormalDriverStrategy{}}}
}

func main() {
	sportsCar := NewSportsVehicle()
	offRoadVehicle := NewOffRoadVehicle()
	goodsTruck := NewGoodsVehicle()

	sportsCar.Drive()      // Output: Sports drive strategy
	offRoadVehicle.Drive() // Output: Normal drive strategy
	goodsTruck.Drive()     // Output: Normal drive strategy
}
