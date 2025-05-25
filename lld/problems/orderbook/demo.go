package main

import (
	"problem/engine"
	"problem/types"
)

func main() {
	ob := engine.NewOrderBook()

	ob.SubmitOrder(&types.Order{Type: types.Buy, Price: 100.0, Quantity: 10})
	ob.SubmitOrder(&types.Order{Type: types.Sell, Price: 99.0, Quantity: 5})
	ob.SubmitOrder(&types.Order{Type: types.Sell, Price: 100.0, Quantity: 5})
	ob.SubmitOrder(&types.Order{Type: types.Buy, Price: 101.0, Quantity: 7})
	ob.SubmitOrder(&types.Order{Type: types.Sell, Price: 98.0, Quantity: 2})
	ob.SubmitOrder(&types.Order{Type: types.Buy, IsMarket: true, Quantity: 4}) // Market order
}
