package model

type VendingMachine struct {
	CurrentState       State
	Items              map[int]Item
	ItemCounts         map[int]int
	SelectedItem       int
	CoinsForSelectItem int
	SelectedCount      int
}

func (vm *VendingMachine) UpdateItem(item Item, count int) {
	// Assumption item and count is validated in controller.
	vm.Items[item.Id] = item
	_, exist := vm.Items[item.Id]
	if exist {
		vm.ItemCounts[item.Id] = vm.ItemCounts[item.Id] + count
	} else {
		vm.ItemCounts[item.Id] = count
	}
}

func (vm *VendingMachine) Start() {
	// Assumption item and count is validated in controller.
	vm.CurrentState.Process(vm)
}
