package main

import "fmt"

// FanState defines the interface for fan states.
type FanState interface {
	PullChain(fan *Fan)
}

// FanOffState is a concrete implementation of FanState for the Off state.
type FanOffState struct{}

func (s *FanOffState) PullChain(fan *Fan) {
	fmt.Println("Turning fan on to low.")
	fan.SetState(&FanLowState{})
}

// FanLowState is a concrete implementation of FanState for the Low state.
type FanLowState struct{}

func (s *FanLowState) PullChain(fan *Fan) {
	fmt.Println("Turning fan on to high.")
	fan.SetState(&FanHighState{})
}

// FanHighState is a concrete implementation of FanState for the High state.
type FanHighState struct{}

func (s *FanHighState) PullChain(fan *Fan) {
	fmt.Println("Turning fan off.")
	fan.SetState(&FanOffState{})
}

// Fan represents the context in which the states operate.
type Fan struct {
	currentState FanState
}

func (f *Fan) SetState(state FanState) {
	f.currentState = state
}

func (f *Fan) PullChain() {
	f.currentState.PullChain(f)
}

func main() {
	fan := &Fan{currentState: &FanOffState{}}
	fan.PullChain() // Turning fan on to low.
	fan.PullChain() // Turning fan on to high.
	fan.PullChain() // Turning fan off.
}
