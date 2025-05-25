package types

import "time"

type OrderType string

const (
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

type Order struct {
	ID        int64
	Type      OrderType
	Price     float64
	Quantity  int
	Timestamp time.Time
	IsMarket  bool
}
