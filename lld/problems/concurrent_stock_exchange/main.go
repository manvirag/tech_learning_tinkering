package main

import (
  "main/models";
	"main/usecases";
	"fmt"
	
)

func main() {

	userA :=   models.User{
		UserId:    "userA",
		UserName:  "Alice",
		AccountId: "accountA",
	}

	userB :=   models.User{
		UserId:    "userB",
		UserName:  "Bob",
		AccountId: "accountB",
	}

	accountA := models.Account{
		AccountId: "accountA",
		Balance:   100,
		Orders:    []models.Order{},
		PortFolio: models.PortFolio{
			PortFolioId: "portFolioA",
			Stocks:      make([]string,0),
			StockCounts: map[string]int{},
		},
		User: userA,
	}

	accountB := models.Account{
		AccountId: "accountB",
		Balance:   200,
		Orders:    []models.Order{},
		PortFolio: models.PortFolio{
			PortFolioId: "portFolioB",
			Stocks:      make([]string,0),
			StockCounts: map[string]int{},
		},
		User: userB,
	}

	accountRepo := models.NewAccountRepo()
	userRepo := models.NewUserRepository(accountRepo)
	userRepo.CreateUser(userA)
	userRepo.CreateUser(userB)
	accountRepo.CreateAccount(accountA)
	accountRepo.CreateAccount(accountB)

	fmt.Println("\n...............................................")
	// user and their account is created
	updatedAcc, _ := accountRepo.GetAccount("accountA")
	fmt.Printf("Initial Account Status: %+v\n\n", updatedAcc);
	


	// sample stock creation
	stockExRepo := models.NewStockExchangeRepository()
    stockA  := models.Stock {	
		StockId: "stockA",
		StockName: "Stock A",
		StockPrice: 50,
	}
	stockExRepo.AddStock(stockA)

	fmt.Println(stockExRepo.GetStocks())
	
	
	stockExchangeService := usecases.NewStockExchangeUsercase(userRepo, accountRepo, stockExRepo, &models.SimpleOrderExecution{})	
    quantity := 2
	stockExchangeService.BuyStock("stockA", "accountA", quantity);
	updatedAccount , _ := accountRepo.GetAccount("accountA")
	fmt.Println("\n...............................................")
	fmt.Printf("updatedAccount after buying stock: %+v,  with quantity: %d:   %+v \n" , stockA, quantity, updatedAccount)


	err := stockExchangeService.BuyStock("stockA", "accountA", 2);
	if err != nil {
		fmt.Println("\n Error occured while buying stocks %+v \n", err)
	}

    quantity = 1
	err = stockExchangeService.SellStock("stockA", "accountA", quantity);
	if err != nil {
		fmt.Println("\nError occured while selling stocks %+v \n", err)
	}


	updatedAccount2 , _ := accountRepo.GetAccount("accountA")
	fmt.Println("\n...............................................")
	fmt.Printf("updatedAccount after buying stock: %+v,  with quantity: %d:   %+v \n" , stockA, quantity, updatedAccount2)
	

	

	
	


	
}
