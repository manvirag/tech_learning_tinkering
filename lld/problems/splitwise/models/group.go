package models

import (
	"errors"
	"fmt"
	"math"
)

type Group struct{
  GroupId string `json:"group_id"`
  Name string  `json:"name"`
  Users []User
  Expenses []Expense
  UserBalance map[User]float64  // negative -> user has to take, pos- > has to give. means user want to make zero their balance
}

func (g Group)CheckUser(userId string) error{
  for _, user := range g.Users{
    if user.UserId == userId{
      return nil
    }
  }
  return errors.New(fmt.Sprintf("User : %s not found", userId))
}
func (g Group)Validate(expense Expense) error {
  toUserExist := true
  FromUserExist := true
    // Add logic checking uservssplit users as well as paid by
  
  if FromUserExist && toUserExist {
    return nil
  }
  
  return errors.New("User doesn't exist in group")
    
}
func (g Group)Balancify(expense Expense) Group{
    transactions := expense.GetTransactions()
     
    if g.UserBalance == nil {
      g.UserBalance = make(map[User]float64)
    }
    for _, transaction := range transactions{
        if transaction.From == transaction.To{
          continue
        }
        for _, user := range g.Users{
          if user.UserId == transaction.From{
            
            if _, ok := g.UserBalance[user]; !ok {
              
              g.UserBalance[user] = 0.0
              
            }
            g.UserBalance[user] = g.UserBalance[user] - transaction.Amount  
            
          } else if (user.UserId == transaction.To) {
            if _, ok := g.UserBalance[user]; !ok {
              g.UserBalance[user] = 0.0
            }
            g.UserBalance[user] = g.UserBalance[user] + transaction.Amount  
          }
        }
    }
     
    return g;
}


// Algorithm to find user transactions, as of now don't focussing on min transactions, focussing on correctness. This will have atmost tx equal to users of group
// Assume this set of transactions will be same for all user , like can use db to store this or cache it
func (g Group)FindMinSplitwiseTransactions() []Transaction{
  splitWiseMinTransactions := make([]Transaction,0)
  negatives := make([]Transaction,0)
  positives := make([]Transaction,0)
  
  for user, balance := range g.UserBalance{
    if balance < 0 {
      negatives = append(negatives, Transaction{
        From:  user.UserId,
        Amount: balance,
      })
    } else if balance > 0 {
      positives = append(positives, Transaction{
        From:  user.UserId,
        Amount: balance,
      })
    }
  }
  
  // fmt.Println(negatives)
  // fmt.Println(positives)
  g.SplitWiseBasicAlgo(&splitWiseMinTransactions, negatives, positives)
  fmt.Printf("Splitwise Transactions : %+v \n", splitWiseMinTransactions)

  return splitWiseMinTransactions
}

func (g Group)SplitWiseBasicAlgo(splitWiseMinTransactions *[]Transaction, negatives []Transaction, positives []Transaction){

  newNegatives := make([]Transaction,0)
  newPositives := make([]Transaction,0)
  // fmt.Println("fdasfasdf")
  // fmt.Println(splitWiseMinTransactions)
  for _, negative := range negatives{
    for _, positive := range positives{
        *splitWiseMinTransactions = append(*splitWiseMinTransactions, Transaction{
          To: negative.From,
          From: positive.From,
          Amount: math.Abs(negative.Amount),
        })
        if len(negatives)>1 {
          newNegatives = negatives[1:]
        }
        positive.Amount = positive.Amount + negative.Amount
        // fmt.Println(positive)
        if positive.Amount < 0 {
          newNegatives = append(newNegatives, positive)
          if len(positives)>1{
             newPositives = positives[1:]
          }
        } else if positive.Amount > 0{
          newPositives = positives
        } 
        g.SplitWiseBasicAlgo(splitWiseMinTransactions, newNegatives, newPositives)
        break;
    }
    break;
  }
  
}