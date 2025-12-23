package models

import (
  "errors"
  "fmt"
   "math"
)

type ExpenseType string

const (
  EXACT      ExpenseType = "EXACT"
  EQUAL      ExpenseType = "EQUAL"
  PERCENTAGE ExpenseType = "PERCENTAGE"
)

type Expense interface {
  Validate() error
  GetTransactions() []Transaction
}

type BasicExpense struct {
  Amount      float64            `json:"amount"`
  Description string             `json:"description"`
  PaidBy      string             `json:"paid_by"`
  UserVsSplit map[string]float64 // exact -> total amount, percentage -> percentage, exact -> total amount
}

func (e BasicExpense) Validate() error {
  if e.Amount <= 0 {
    return errors.New("amount must be greater than zero")
  }
  if e.PaidBy == "" {
    return errors.New("paid_by cannot be empty")
  }
  if len(e.UserVsSplit) == 0 {
    return errors.New("user_vs_split cannot be empty")
  }
  for user, split := range e.UserVsSplit {
    if user == "" {
      return errors.New("user in user_vs_split cannot be empty")
    }
    if split <= 0 {
      return fmt.Errorf("split for user %s must be greater than zero", user)
    }
  }
  return nil
}
func (e  BasicExpense) GetTrasactions() []Transaction { return []Transaction{}}

type EqualExpense struct {
  BasicExpense
  Type ExpenseType `json:"type"`
}

func (e EqualExpense) Validate() error {
  if err := e.BasicExpense.Validate(); err != nil {
    return err
  }
  if e.Type != EQUAL {
    return errors.New("expense type must be EQUAL")
  }

  expectedSplit := math.Round(e.Amount/float64(len(e.UserVsSplit))*100) / 100

  fmt.Println(expectedSplit)
  for _, split := range e.UserVsSplit {
    if split != expectedSplit {
      return errors.New("splits are not equal for an EQUAL expense")
    }
  }
  return nil
}

func (e EqualExpense)GetTransactions() []Transaction {
  transactions := make([]Transaction, 0, len(e.UserVsSplit))
  for user, split := range e.UserVsSplit{
    transactions = append(transactions, Transaction{
      From: e.PaidBy,
      To: user,
      Amount: split,
    })
  }
  return transactions
}

type ExactExpense struct {
  BasicExpense
  Type ExpenseType `json:"type"`
}

func (e ExactExpense) Validate() error {
  if err := e.BasicExpense.Validate(); err != nil {
    return err
  }
  if e.Type != EXACT {
    return errors.New("expense type must be EXACT")
  }

  totalSplit := 0.0
  for _, split := range e.UserVsSplit {
    totalSplit += split
  }
  if totalSplit != e.Amount {
    return errors.New("total of splits does not equal the amount for an EXACT expense")
  }
  return nil
}

func (e ExactExpense)GetTransactions() []Transaction {
  transactions := make([]Transaction, 0, len(e.UserVsSplit))
  for user, split := range e.UserVsSplit{
    transactions = append(transactions, Transaction{
      From: e.PaidBy,
      To: user,
      Amount: split,
    })
  }
  return transactions
}

type PercentageExpense struct {
  BasicExpense
  Type ExpenseType `json:"type"`
}

func (e PercentageExpense) Validate() error {
  if err := e.BasicExpense.Validate(); err != nil {
    return err
  }
  if e.Type != PERCENTAGE {
    return errors.New("expense type must be PERCENTAGE")
  }
  totalPercentage := 0.0
  for _, percentage := range e.UserVsSplit {
    totalPercentage += percentage
  }
  if totalPercentage != 100.0 {
    return errors.New("total of percentages does not equal 100% for a PERCENTAGE expense")
  }
  return nil
}

func (e PercentageExpense)GetTransactions() []Transaction {
  transactions := make([]Transaction, 0, len(e.UserVsSplit))
  for user, split := range e.UserVsSplit{
    amount := math.Round((e.Amount * split)/100)
    transactions = append(transactions, Transaction{
      From: e.PaidBy,
      To: user,
      Amount: amount,
    })
  }
  return transactions
}






