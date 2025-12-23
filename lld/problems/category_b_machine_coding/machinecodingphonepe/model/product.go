package model

type Product struct {
	ID   string
	Name string
}

type ProductUsecase interface {
	GetProduct(productId string) (*Product, error)
}

type ProductRepository interface {
	GetProductByID(productId string) (*Product, error)
}
