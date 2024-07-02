package model

import (
	"fmt"
)

type VerificationState struct{}

func (is *VerificationState) Process(machine *VendingMachine) {
	fmt.Println("Verifying....Details")
	if item, exist := machine.Items[machine.SelectedItem]; exist && machine.ItemCounts[machine.SelectedItem] >= machine.SelectedCount && item.Price*machine.SelectedCount <= machine.CoinsForSelectItem {
		machine.CurrentState = &OutputState{}
		machine.CurrentState.Process(machine)
	} else {
		fmt.Println("Verification failed. Either we have less quantity of item or item doesn't exist . Please contact to support")
		machine.CurrentState = &IdleState{}
	}

}
