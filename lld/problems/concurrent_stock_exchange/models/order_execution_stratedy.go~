package models

// import "fmt"

type OrderExecutionStratedy interface {
  BuyStock(stock Stock, account Account, quantity int) error
  SellStock(stock Stock, account Account, quantity int) error
}

// Assume there is always seller and there is always buyer
type SimpleOrderExecution struct{
  
}

func (soe *SimpleOrderExecution )BuyStock(stock Stock, account Account, quantity int) error{
  // fmt.Printf("User %s , buying stocks: %s, with price %.2f \n", account.User.UserName ,stock.StockName, stock.StockPrice)
  return nil
}

func (soe *SimpleOrderExecution) SellStock(stock Stock, account Account, quantity int) error{
  // fmt.Printf("User %s , selling stocks: %s, with price %.2f \n", account.User.UserName , stock.StockName, stock.StockPrice )
  return nil
}