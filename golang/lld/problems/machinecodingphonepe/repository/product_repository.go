package repository

import (
	"errors"
	"main/model"
)

type InMemoryProductRepository struct {
	products map[string]model.Product
}

var inMemoryProductRepository *InMemoryProductRepository

func NewInMemoryProductRepository() *InMemoryProductRepository {
	if inMemoryProductRepository == nil {
		inMemoryProductRepository = &InMemoryProductRepository{
			products: map[string]model.Product{
				"1": {ID: "1", Name: "Product 1"},
				"2": {ID: "2", Name: "Product 2"},
				"3": {ID: "3", Name: "Product 3"},
			},
		}
	}
	return inMemoryProductRepository
}

func (r *InMemoryProductRepository) GetProductByID(productID string) (*model.Product, error) {
	product, exists := r.products[productID]
	if !exists {
		return nil, errors.New("Product not found")
	}
	return &product, nil
}
