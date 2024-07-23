package main

import (
	"fmt"
	"main/models"
	"main/usecases"
	"strconv"
	"math"
)

func main() {
	userRepo :=	models.NewUserRepo()
	for i:=0;i<5;i++{
			user := models.User{UserId: strconv.Itoa(i) , Name: string((65+i))}
			// userJson, _ := json.Marshal(user)
		 //  fmt.Println(string(userJson))
		  userRepo.AddUSer(user)
	}
	// fmt.Println(userRepo)

	splitwise := usecases.NewSplitWise(userRepo)

	
	group := models.Group{GroupId: "1", Name: "Group 1", Users: []models.User{userRepo.GetUser("0"), userRepo.GetUser("1"), userRepo.GetUser("2")}}
	splitwise.AddGroup(group)

	exactExpense := models.ExactExpense{
		BasicExpense: models.BasicExpense{
			Amount:       100,
			Description:  "Dinner",
			PaidBy:       "0",
			UserVsSplit:  map[string]float64{"0": 50, "1": 25, "2": 25},
		},
		Type: models.EXACT,
	}
	amount := math.Round(50/float64(3)*100) / 100
	equalExpense := models.EqualExpense{
		BasicExpense: models.BasicExpense{
			Amount:       50,
			Description:  "Coffee",
			PaidBy:       "0",
			UserVsSplit:  map[string]float64{"0": amount, "1": amount, "2": amount},
		},
		Type: models.EQUAL,
	}
	percentageExpense := models.PercentageExpense{
		BasicExpense: models.BasicExpense{
			Amount:       50,
			Description:  "Coffee",
			PaidBy:       "0",
			UserVsSplit:  map[string]float64{"0": 25, "1": 25, "2": 50},
		},
		Type: models.PERCENTAGE,
	}

	err := splitwise.AddExpense(exactExpense, "1")
	if ( err != nil ){
		fmt.Println(err)
	}
	err = splitwise.AddExpense(equalExpense, "1")
	if ( err != nil ){
		fmt.Println(err)
	}
	err = splitwise.AddExpense(percentageExpense, "1")
	if ( err != nil ){
		fmt.Println(err)
	}
	// gr , _ := splitwise.GetGroup("1")
	// splitwiseJson,_  := json.Marshal(gr)
	// fmt.Println(string(splitwiseJson))
	aa, err := splitwise.ShowBalance("1", "0")
	if( err != nil ){
		fmt.Println(err)
	}
	fmt.Printf("Transactions Involving User: 0 , Are: %+v", aa)
}
