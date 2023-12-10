package main

import (
	"fmt"
	"main/controller"
	"main/repository"
	"main/usecases"
)

// Testing code
func main() {
	userRepository := repository.NewInMemoryUserRepository()
	userUsecase := usecases.GetUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	err := userController.CreateUser("John Doe", "john.doe@example.com", "securepassword")
	if err != nil {
		fmt.Printf("Error creating user: %s\n", err)
	} else {
		fmt.Println("User created successfully")
	}

	// fmt.Println(userController.GetUser("123"))

	// ....................
	// productUsecase := usecases.NewProductUsecase(repository.NewInMemoryProductRepository())
	// fmt.Println(productUsecase.GetProduct("1"))

	// ....................

	shoppingCartUsecase := usecases.NewShoppingCartUsecase(userUsecase, usecases.NewProductUsecase(repository.NewInMemoryProductRepository()), repository.NewInMemoryShoppingCartRepository())

	// Add products to the cart
	userID := "user123"
	productID1 := "1"
	productID2 := "2"

	err = shoppingCartUsecase.AddToCart(userID, productID1, 2)
	if err != nil {
		fmt.Printf("Error adding product %s to the cart: %s\n", productID1, err)
	}

	err = shoppingCartUsecase.AddToCart(userID, productID2, 1)
	if err != nil {
		fmt.Printf("Error adding product %s to the cart: %s\n", productID2, err)
	}

	fmt.Printf("\nGetting the shopping cart for user %s...\n", userID)
	cart, err := shoppingCartUsecase.GetCart(userID)
	if err != nil {
		fmt.Printf("Error getting the shopping cart: %s\n", err)
	} else {
		fmt.Printf("Shopping Cart for user %s:\n", userID)
		for _, item := range cart.CartItems {
			fmt.Printf("ProductID: %s, Quantity: %d\n", item.ProductId, item.Quantity)
		}
	}

	// Checkout
	fmt.Printf("\nChecking out the shopping cart for user %s...\n", userID)
	order, err := shoppingCartUsecase.Checkout(userID)
	if err != nil {
		fmt.Printf("Error checking out the shopping cart: %s\n", err)
	} else {
		fmt.Printf("Order Details:\n")
		fmt.Printf("OrderID: %s\n", order.OrderID)
		fmt.Printf("UserID: %s\n", order.UserID)
		fmt.Printf("Items:\n")
		for _, item := range order.OrderedItems {
			fmt.Printf("ProductID: %s, Quantity: %d\n", item.ProductId, item.Quantity)
		}
	}

}
