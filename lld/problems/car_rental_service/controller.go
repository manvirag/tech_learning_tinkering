package main

import (
	"fmt"
	"main/models"
	"main/usecase"
)

func main() {
	user := models.Rider{
		User: models.User{
				Name: "anubhav",
					UserId: "1234",
		},
	}

	driver1 := models.Driver{
		User: models.User{
			Name: "driver1",
			UserId: "d1",
		},
		Status: "Available",
	}

		driver2 := models.Driver{
			User: models.User{
				Name: "driver2",
					UserId: "d12",
			},
			Status: "Available",
	}


		driver3 := models.Driver{
			User: models.User{
					Name: "driver3",
					UserId: "d123",
			},
			Status: "Available",
	}

	driver4 := models.Driver{
		User: models.User{
			Name: "driver4",
			UserId: "d1234",
		},
		Status: "Available",
	}

	userRepo := models.NewUserRepository()
	userRepo.AddUser(user)
	userRepo.AddUser(driver1)
	userRepo.AddUser(driver2)
	userRepo.AddUser(driver3)
	userRepo.AddUser(driver4)

	fmt.Println(userRepo)

	userAccount := models.Account{
		AccountId: "user_1234",
		UserId: "1234",
		Amount: 100,
	}

	accountRepo := models.NewAccountRepo()
	accountRepo.AddAccount(userAccount)

	fmt.Print(accountRepo)

	uber_system := usecase.NewUber(accountRepo , userRepo , usecase.DummyFare{}, usecase.DummyDriverMatchingStrategy{})

	

	ride, errr := uber_system.RequestRide("source", "destination", "1234");
	if(errr!=nil){
		fmt.Println(errr)
	}
	fmt.Println(ride)
		
	
}

