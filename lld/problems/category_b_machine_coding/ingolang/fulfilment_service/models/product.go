package models

type Product struct {
	Id int
}

type ProductInventory struct {
	ProductDetail Product
	Quantity      int
}
