package model

import (
	"fmt"
)

type DisplayState struct{}

func (is *DisplayState) Process(machine *VendingMachine) {
	fmt.Println("Enter the coins")
	fmt.Scanln(&machine.CoinsForSelectItem)
	fmt.Println("Please enter the item id")
	fmt.Scanln(&machine.SelectedItem)
	fmt.Println("Enter the item counts")
	fmt.Scanln(&machine.SelectedCount)

	machine.CurrentState = &VerificationState{}
	machine.CurrentState.Process(machine)
}
