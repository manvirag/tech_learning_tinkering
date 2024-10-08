package models

import (
	"errors"
	"fmt"
)

// account id , money

type BankRepoInterface interface{
  GetBalance(accountId string) (int, error)
  CreateAccount(accountId string, initialAmount int32) error
  Update(accountId string, amount int32) error
  CreateAtm(atmId string, pin string) error
  ValidateAtm(accountId string, pin string) error
}

type BankRepo struct{
 bank map[string]int32;  // account amount
 atm map[string]string;  // account pin
}

func NewBankRepo() *BankRepo{
  return &BankRepo{
    bank: make(map[string]int32),
    atm: make(map[string]string),
  }
}

func (rp *BankRepo) GetBalance(accountId string) (int, error){
  if balance, ok := rp.bank[accountId]; ok{
    return int(balance), nil
  }
  return 0, errors.New("account not found")
}

func (rp *BankRepo) CreateAccount(accountId string, initialAmount int32) error{
    if _ , ok := rp.bank[accountId]; ok {
      return errors.New("Account already exist")
    }

    rp.bank[accountId] = initialAmount

    return nil;
  
}

func (rp *BankRepo) Update(accountId string, amount int32) error{
  if _, ok := rp.bank[accountId]; !ok {
    return errors.New("account not found")
  }

  rp.bank[accountId] += amount
  return nil;      
}

func (rp *BankRepo) CreateAtm(atmId string, pin string) error{
  if _, ok := rp.atm[atmId]; ok {
    return errors.New("atm already exist")
  }
  rp.atm[atmId] = pin
  fmt.Println(rp.atm)
  return nil;
}

func (rp *BankRepo) ValidateAtm(accountId string, pin string) error {
  if _, ok := rp.atm[accountId]; !ok {
    return errors.New("atm not found")
  }

  if rp.atm[accountId] != pin {
    return errors.New("invalid pin")
  }
  return nil
}
