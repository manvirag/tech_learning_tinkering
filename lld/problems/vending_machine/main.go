package main

import (
	"fmt"
	"vending_machine/controller"
	"vending_machine/model"
)

func main() {
	controler := controller.VendingMachineController{
		VendingMachine: &model.VendingMachine{
			CurrentState:       &model.IdleState{},
			CoinsForSelectItem: -1,
			SelectedItem:       -1,
			SelectedCount:      -1,
			Items:              make(map[int]model.Item),
			ItemCounts:         make(map[int]int),
		},
	}

	controler.UpdateItem(model.Item{Id: 12, Name: "Coca", Price: 30}, 5)
	controler.UpdateItem(model.Item{Id: 13, Name: "Pepsi", Price: 40}, 2)
	controler.UpdateItem(model.Item{Id: 3, Name: "Pancake", Price: 10}, 2)
	for {
		fmt.Println("....................Start Vending Machine..................")
		controler.RunVendingMaching()

	}

}
