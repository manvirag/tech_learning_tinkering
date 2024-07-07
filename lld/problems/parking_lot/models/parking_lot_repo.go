package models

import (
	"errors"
	"fmt"
    "sync"
)

type ParkingLotRepoInterface interface {
  AddNewLevel(level int32) error
  GetLevel() int32
  GetSlots(level int32) int32
  Park(vehicle VehicleInterface) (int32, int32, error)
  Release(level int32, slot int32, vehicle VehicleInterface) error  
}

type ParkingLotRepo struct {
    Level int32
    SlotCounts map[int32]int32   // Each level total slots
    SlotsAvailability map[int32]map[int32]bool  // level slot availability
    ParkingSlots map[int32]map[int32]VehicleInterface  // level , slot, vehicle
    sync.RWMutex
}

func NewParkingLotRepo() *ParkingLotRepo{
    return &ParkingLotRepo{
        Level: 0,
        SlotCounts: make(map[int32]int32),
        ParkingSlots: make(map[int32]map[int32]VehicleInterface ),
        SlotsAvailability: make(map[int32]map[int32]bool),
    }
}

func (p *ParkingLotRepo) AddNewLevel(slots int32) error {
    defer p.Unlock()
    p.Lock()
    p.Level = p.Level + 1
    p.SlotCounts[p.Level] = slots
    p.SlotsAvailability[p.Level] = make(map[int32]bool)
    for i := 1; i <= int(slots); i++{
        p.SlotsAvailability[p.Level][int32(i)] = false
    }
    p.ParkingSlots[p.Level] = make(map[int32]VehicleInterface)
    for i := 1; i <= int(slots); i++{
         p.ParkingSlots[p.Level][int32(i)] = nil
    }
    return nil
}

func (p *ParkingLotRepo) GetLevel() int {
    p.RLock()
    defer p.RUnlock()
    return int(p.Level)
}


func (p *ParkingLotRepo) GetSlots(level int32) int32 {
    p.RLock()
    defer p.RUnlock()
    return p.SlotCounts[level]
}

func (p *ParkingLotRepo) Park(vehicle VehicleInterface) (int32, int32, error) {
    p.Lock()
    defer p.Unlock()
    for i := 1; i <= int(p.Level); i++ {
        if p.SlotCounts[int32(i)] > 0 {
            for slot := 1; slot <= int(p.SlotCounts[int32(i)]); slot++ {
                if !p.SlotsAvailability[int32(i)][int32(slot)] {
                    p.SlotsAvailability[int32(i)][int32(slot)] = true
                    p.SlotCounts[int32(i)]--
                    if _, ok := p.ParkingSlots[int32(i)]; !ok {
                        p.ParkingSlots[int32(i)] = make(map[int32]VehicleInterface)
                    }
                    p.ParkingSlots[int32(i)][int32(slot)] = vehicle
                    return int32(i), int32(slot), nil
                }
            }
        }
    }
    return int32(0),int32(0), errors.New("no available slot")
}


func(p *ParkingLotRepo ) Release(level int32, slot int32, vehicle VehicleInterface) error {
    p.Lock()
    defer p.Unlock()
    if(p.Level < level || p.SlotCounts[level] < slot) {
        return errors.New("Invalid Level and slot number");
    }

    if !p.SlotsAvailability[level][slot] {
        return errors.New("This slot is empty , please give correct slot and level")
    }

    veh := p.ParkingSlots[level][slot]

    if(veh.GetVehicleNumber() != vehicle.GetVehicleNumber()){
        return errors.New(fmt.Sprintf("This slot is not occupied by this vehicle: %+v \n", veh))
    }

    p.SlotCounts[level]++;
    p.SlotsAvailability[level][slot] = true
    p.ParkingSlots[level][slot] = nil
    return nil;
}