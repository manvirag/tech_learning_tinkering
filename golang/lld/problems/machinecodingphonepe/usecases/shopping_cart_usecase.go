package usecases

import (
	"errors"
	"main/model"
)

type ShoppingCartUsecaseImpl struct {
	userUsecase            model.UserUsecase
	productUsecase         model.ProductUsecase
	shoppingCartRepository model.ShoppingCartRepository
}

var shoppingCartUsecaseImplVar *ShoppingCartUsecaseImpl

func NewShoppingCartUsecase(userUsecase model.UserUsecase, productUsecase model.ProductUsecase, shoppingCartRepository model.ShoppingCartRepository) *ShoppingCartUsecaseImpl {
	if shoppingCartUsecaseImplVar == nil {
		shoppingCartUsecaseImplVar = &ShoppingCartUsecaseImpl{
			userUsecase:            userUsecase,
			productUsecase:         productUsecase,
			shoppingCartRepository: shoppingCartRepository,
		}
	}
	return shoppingCartUsecaseImplVar
}

func (uc *ShoppingCartUsecaseImpl) AddToCart(userID, productID string, quantity int) error {

	if userID == "" || productID == "" {
		return errors.New("Invalid userId or productId")
	}

	_, err := uc.userUsecase.GetUser(userID)
	if err != nil {
		return errors.New("user not found")
	}

	_, errProduct := uc.productUsecase.GetProduct(productID)
	if errProduct != nil {
		return errors.New("product not found")
	}

	return uc.shoppingCartRepository.AddToCart(userID, productID, quantity)
}

func (uc *ShoppingCartUsecaseImpl) GetCart(userID string) (*model.ShoppingCart, error) {
	_, err := uc.userUsecase.GetUser(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return uc.shoppingCartRepository.GetCart(userID)
}

func (uc *ShoppingCartUsecaseImpl) Checkout(userID string) (*model.Order, error) {
	_, err := uc.userUsecase.GetUser(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return uc.shoppingCartRepository.Checkout(userID)
}
