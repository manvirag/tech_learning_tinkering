package main

import (
	"main/models"
	"fmt"
)

func main() {

	  bankRepo := models.NewBankRepo()
	  user := models.User{
			UserId: "user_A",
			AccountId: "account_A",
		}
	  bankRepo.CreateAccount(user.AccountId, 10)
	  amount , err := bankRepo.GetBalance(user.AccountId)
	  if(err != nil ){
			fmt.Println(err)
		}
	  fmt.Println(amount)
	  bankRepo.CreateAtm("account_A", "1234")
	  

	  atm := models.NewAtm(bankRepo)

		for {
			atm.Process()
		}
	  
}
