package usecases

import (
	
	"errors"
	"fmt"
	"main/models"
)

type SplitWiseInterface interface{
  AddExpense(expense models.Expense, groupId string) error
  AddGroup(group models.Group) error
  GetGroup(groupId string) (models.Group, error)
  ShowBalanace(groupId string, userId string) ([]models.Transaction, error)
}


type SplitWise struct{
  userRepo *models.UserRepo
  groups map[string]models.Group  `json:"groups"`  // ideally should be a different group repo itself.
  
  
}

func NewSplitWise(userRepo *models.UserRepo) *SplitWise{
  return &SplitWise{
    userRepo: userRepo,
    groups: make(map[string]models.Group),
  }
}

func (s *SplitWise) AddExpense(expense models.Expense, groupId string) error {
  
  err := expense.Validate()
  if err != nil {
    return err
  }
  
	if group, ok := s.groups[groupId]; ok {
    err = group.Validate(expense)
    if err != nil {
      return err;
    }
    group.Expenses = append(group.Expenses, expense)
    group = group.Balancify(expense)
    fmt.Println(group)
    s.groups[groupId] = group
    return nil
	}
  return errors.New( fmt.Sprintf("Group : %s not found", groupId))
}

func (s *SplitWise) AddGroup(group models.Group) error {
  if(len(group.Users) ==0 ){
    return errors.New("Group must have atleast one user")
  }
	s.groups[group.GroupId] = group
	return nil
}

func (s *SplitWise)GetGroup(groupId string) (models.Group, error){
  if group, ok := s.groups[groupId]; ok {
    return group, nil
  }
  return models.Group{}, errors.New( fmt.Sprintf("Group : %s not found", groupId))
}

func (s *SplitWise)ShowBalance(groupId string, userId string) ([]models.Transaction, error) {
  splitWiseMinTransactions := make([]models.Transaction,0)
  if group, ok := s.groups[groupId]; ok {
    err := group.CheckUser(userId)
    if err != nil {
      return nil, err
    }
    splitWiseMinTransactions = group.FindMinSplitwiseTransactions()
    fmt.Println(group)
  }  
  
  fmt.Printf("\nSplitwise Transactions : %+v \n", splitWiseMinTransactions)

  

  transactionsInvolvingUser := []models.Transaction{}

  for _, splitWiseMinTransaction := range splitWiseMinTransactions {
    if splitWiseMinTransaction.From == userId || splitWiseMinTransaction.To == userId {
    transactionsInvolvingUser = append(transactionsInvolvingUser, splitWiseMinTransaction)
    }
  }
  
  return transactionsInvolvingUser, nil
}