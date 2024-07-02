package model

type IdleState struct{}

func (is *IdleState) Process(machine *VendingMachine) {
	machine.CurrentState = &DisplayState{}
	machine.CurrentState.Process(machine)
}
