package models

type OrderTypeEnum string

const (
	BUY  OrderTypeEnum = "BUY"
	SELL OrderTypeEnum = "SELL"
)


type Order struct{
  OrderId string
  AccountId string
  StockId string
  Quantity int
  Price float64 
  OrderType OrderTypeEnum
}