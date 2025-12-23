package controller

import (
	"vending_machine/model"
)

type VendingMachineController struct {
	VendingMachine *model.VendingMachine
}

func (v *VendingMachineController) RunVendingMaching() { // making independent of user
	// request validation
	// get item from vending machine
	v.VendingMachine.Start()
	// return response
}

func (v *VendingMachineController) ShowAllItems() error {
	return nil
}

func (v *VendingMachineController) DeleteItem(itemId int) error {
	// Admin Access
	return nil
}

func (v *VendingMachineController) UpdateItem(item model.Item, frequency int) {
	// Admin Access
	v.VendingMachine.UpdateItem(item, frequency)
}
