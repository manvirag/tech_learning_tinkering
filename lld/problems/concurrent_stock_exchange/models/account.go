package models


import  ( "errors"; "fmt" )

type Account struct{
  AccountId string 
  Balance float64
  Orders []Order
  PortFolio PortFolio 
  User User
}

type AccountRepo struct {
  accounts map[string]Account
}

type AccountRepoInterface interface{
  GetAccount(accountId string) (Account, error)
  CreateAccount(account Account) 
  UpdateAccount(account Account)
  DeleteAccount(accountId string)
  GetPorfolio(accountId string) PortFolio
  AddOrder(accountId string, order Order) error
}

func NewAccountRepo() *AccountRepo{
  return &AccountRepo{
    accounts: make(map[string]Account),
  }
}

func (ar *AccountRepo) CreateAccount(account Account){
    ar.accounts[account.AccountId] = account
}

func (ar *AccountRepo) GetAccount(accountId string) ( Account, error ){
  for id, account := range ar.accounts {
    if (id == accountId) {
      return account, nil
    }
  }

  return Account{}, errors.New("Account not found")
}

func (ar *AccountRepo ) AddOrder(accountId string, order Order) error {
   if account , ok := ar.accounts[accountId]; ok {
     account.Orders = append(account.Orders, order) 
     if ( order.OrderType == BUY ) {
        account.Balance = account.Balance - order.Price * float64(order.Quantity)
     } else {
        account.Balance = account.Balance + order.Price * float64(order.Quantity)
     }
     portFolio := account.PortFolio

     if counts , exist := portFolio.StockCounts[order.StockId]; exist{
        counts = counts + order.Quantity
        portFolio.StockCounts[order.StockId] = counts
     } else {
        portFolio.StockCounts[order.StockId] = order.Quantity
        portFolio.Stocks = append(portFolio.Stocks, order.StockId)
        
     }

     account.PortFolio = portFolio

     ar.accounts[accountId] = account
     
   }
   
       
  fmt.Printf("Order Log: %+v",ar.accounts[accountId].Orders)
  return nil
}

func (ar *AccountRepo) HasStock(stockId string, accountId string, quantity int) bool {
  if account, ok := ar.accounts[accountId]; ok {
    if counts, exist := account.PortFolio.StockCounts[stockId]; exist {
      return counts >= quantity
    }
  }
  return false
}