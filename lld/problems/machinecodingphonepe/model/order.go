package model

type Order struct {
	OrderID      string
	UserID       string
	OrderedItems []CartItem
}
