package dao

import "fulfilment/models"

type IProductDao interface {
	AddInventory(productId int, amount int) error
	RemoveInventory(productId int) error
	ViewInventory() []models.ProductInventory
}

type ProductDao struct {
	ProductData map[int]*models.ProductInventory
}

func NewProductDao() IProductDao {
	return &ProductDao{ProductData: make(map[int]*models.ProductInventory)}
}

func (pd *ProductDao) AddInventory(productId int, amount int) error {
	return nil
}

func (pd *ProductDao) RemoveInventory(productId int) error {
	return nil
}

func (pd *ProductDao) ViewInventory() []models.ProductInventory {
	result := []models.ProductInventory{}
	return result
}
