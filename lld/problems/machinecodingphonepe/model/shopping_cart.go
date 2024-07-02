package model

type CartItem struct {
	ProductId string
	Quantity  int
}
type ShoppingCart struct {
	CartItems []CartItem
}

type ShoppingCartUsecase interface {
	AddToCart(userID, productID string, quantity int) error
	GetCart(userID string) (*ShoppingCart, error)
	Checkout(userID string) (*Order, error)
}

type ShoppingCartRepository interface {
	AddToCart(userID, productID string, quantity int) error
	GetCart(userID string) (*ShoppingCart, error)
	Checkout(userID string) (*Order, error)
}
