package models

type Account struct{
  AccountId string
  UserId string
  Amount float32
}

type AccountRepo  struct{
  Accounts map[string]Account
}

func NewAccountRepo() *AccountRepo{
  return &AccountRepo{
    Accounts: make(map[string]Account),
  }
}

func (repo *AccountRepo) AddAccount(account Account){
  repo.Accounts[account.AccountId] = account
}

func (repo *AccountRepo) GetAccount(accountId string) Account{
  return repo.Accounts[accountId]
}