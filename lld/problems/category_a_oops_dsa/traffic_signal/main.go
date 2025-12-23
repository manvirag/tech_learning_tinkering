package main

import (
	"fmt"
	"time"
)

type SignalType int

const (
	Red SignalType = iota
	Yellow
	Green
)

func (s SignalType) String() string {
	return [...]string{"Red", "Yellow", "Green"}[s]
}

type SignalConfig struct {
	RedDuration    int
	YellowDuration int
	GreenDuration  int
}

var defaultConfig = SignalConfig{
	RedDuration:    3,
	YellowDuration: 5,
	GreenDuration:  2,
}

type TrafficController struct {
	Config SignalConfig
	Signal SignalType
}

func NewTrafficController(config SignalConfig) *TrafficController {
	return &TrafficController{
		Config: config,
		Signal: Red,
	}
}

func (tc *TrafficController) Start() {
	for {
		tc.runSignal(Red, tc.Config.RedDuration)
		tc.runSignal(Green, tc.Config.GreenDuration)
		tc.runSignal(Yellow, tc.Config.YellowDuration)
	}
}

func (tc *TrafficController) runSignal(signal SignalType, duration int) {
	tc.Signal = signal
	fmt.Printf("Signal: %s\n", signal)
	time.Sleep(time.Duration(duration) * time.Second)
}

func (tc *TrafficController) HandleEmergency() {
	fmt.Println("Emergency detected! Switching to green signal for emergency vehicle.")
	tc.runSignal(Green, tc.Config.GreenDuration)
}

func main() {
	tc := NewTrafficController(defaultConfig)
	go tc.Start()

	// Simulate an emergency after 10 seconds
	time.Sleep(10 * time.Second)
	tc.HandleEmergency()
}
