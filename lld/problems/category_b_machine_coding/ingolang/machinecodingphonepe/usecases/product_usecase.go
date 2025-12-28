package usecases

import (
	"errors"
	"main/model"
)

type ProductUsecase struct {
	productRepository model.ProductRepository
}

func NewProductUsecase(productRepository model.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepository: productRepository,
	}
}

func (c *ProductUsecase) GetProduct(productID string) (*model.Product, error) {
	if productID == "" {
		return nil, errors.New("invalid product id")
	}

	return c.productRepository.GetProductByID(productID)
}
