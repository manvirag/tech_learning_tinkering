package repository

import (
	"errors"
	"fmt"
	"main/model"
	"strconv"
	"time"
)

type InMemoryShoppingCartRepository struct {
	carts map[string]model.ShoppingCart
}

var inMemoryShoppingCartRepositoryVar *InMemoryShoppingCartRepository

func NewInMemoryShoppingCartRepository() *InMemoryShoppingCartRepository {
	if inMemoryShoppingCartRepositoryVar == nil {
		inMemoryShoppingCartRepositoryVar = &InMemoryShoppingCartRepository{
			carts: make(map[string]model.ShoppingCart),
		}
	}

	return inMemoryShoppingCartRepositoryVar
}

func (r *InMemoryShoppingCartRepository) AddToCart(userID, productID string, quantity int) error {
	cart, exists := r.carts[userID]
	if !exists {
		cart = model.ShoppingCart{CartItems: []model.CartItem{}}
		r.carts[userID] = cart
	}

	for _, item := range cart.CartItems {
		if item.ProductId == productID {
			item.Quantity += quantity
			return nil
		}
	}

	cart.CartItems = append(cart.CartItems, model.CartItem{ProductId: productID, Quantity: quantity})
  r.carts[userID] = cart

	fmt.Printf("Updated CartItems: %+v", r.carts)

	return nil
}

func (r *InMemoryShoppingCartRepository) GetCart(userID string) (*model.ShoppingCart, error) {
	cart, exists := r.carts[userID]
	if !exists {
		return &model.ShoppingCart{}, nil
	}
	return &cart, nil
}

func (r *InMemoryShoppingCartRepository) Checkout(userID string) (*model.Order, error) {

	fmt.Println(r.carts)
	cart, exists := r.carts[userID]
	if !exists || len(cart.CartItems) == 0 {
		return nil, errors.New("empty cart")
	}

	order := &model.Order{
		OrderID:      strconv.Itoa(int(time.Now().Unix())),
		UserID:       userID,
		OrderedItems: cart.CartItems,
	}

	delete(r.carts, userID)

	return order, nil
}
