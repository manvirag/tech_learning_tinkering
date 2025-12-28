package model

type State interface {
	Process(vendingMachine *VendingMachine)
}
