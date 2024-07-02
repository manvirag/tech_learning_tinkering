package model

import (
	"fmt"
)

type OutputState struct{}

func (is *OutputState) Process(machine *VendingMachine) {
	fmt.Println("Verification Complete. Please collect your items")
	fmt.Printf("Item: %v \n", machine.Items[machine.SelectedItem].Name)
	fmt.Printf("Item Counts: %v \n", machine.SelectedCount)
	fmt.Println("Have a Nice Day. Please Try Again Thankss.")

	// update
	machine.ItemCounts[machine.SelectedItem] = machine.ItemCounts[machine.SelectedItem] - machine.SelectedCount
	if machine.ItemCounts[machine.SelectedCount] == 0 {
		delete(machine.Items, machine.SelectedItem)
		delete(machine.ItemCounts, machine.SelectedItem)
	}

	machine.CurrentState = &IdleState{}
}
