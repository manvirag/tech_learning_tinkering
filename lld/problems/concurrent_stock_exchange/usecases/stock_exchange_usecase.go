package usecases

import ("main/models";
       "time";
        "fmt";

        "errors"
       )

type StockExchangeInterface interface {
  // All operation which we are doing right now direct with repository.
  BuyStock(stockId string, accountId string, quantity int) error
  SellStock(stockId string, accountId string, quantity int) error
  
}

type StockExchangeUsercase struct {
	userRepo    *models.UserRepository
	accountRepo *models.AccountRepo
	stockRepo   *models.StockExchangeRepository
  orderExecuter models.OrderExecutionStratedy
}

func NewStockExchangeUsercase(userRepo *models.UserRepository, accountRepo *models.AccountRepo, stockRepo *models.StockExchangeRepository, orderExecution models.OrderExecutionStratedy) *StockExchangeUsercase {
	return &StockExchangeUsercase{
		userRepo:    userRepo,
		accountRepo: accountRepo,
		stockRepo:   stockRepo,
    orderExecuter: orderExecution, 
	}
}


func (seu * StockExchangeUsercase) BuyStock(stockId string, accountId string, quantity int) error {
  // get stock from stockRepo

  stock, err := seu.stockRepo.GetStock(stockId)

  if ( err != nil ){
    fmt.Println("Error occured while getting stock %+v", err)
    return err
  }
  
  // get account from accountRepo
  account, err := seu.accountRepo.GetAccount(accountId) 

  if ( err != nil ){
    fmt.Println("Error occured while getting account %+v", err)
    return err
  }
  
  
  // check balance
  if account.Balance < float64(quantity) * stock.StockPrice {
    return errors.New("Insufficient balance")
  }

  err = seu.orderExecuter.BuyStock(stock,account,quantity)
  if err != nil {
    return err
  }
  // buy stock with order execution
  order := models.Order{
    StockId:     stockId,
    AccountId:   accountId,
    Quantity:     quantity,
    Price:        stock.StockPrice,
    OrderId:      time.Now().String(),
    OrderType:    models.BUY,
  }

  seu.accountRepo.AddOrder(accountId, order)



  
  return nil
}




func (seu *StockExchangeUsercase) SellStock(stockId string, accountId string, quantity int) error {
  // get stock from stockRepo
  stock, err := seu.stockRepo.GetStock(stockId)

  if ( err != nil ){
    fmt.Println("Error occured while getting stock %+v", err)
    return err
  }
  
  // get account from accountRepo
  account, err := seu.accountRepo.GetAccount(accountId) 

  if ( err != nil ){
    fmt.Println("Error occured while getting account %+v", err)
    return err
  }
  
  // check balance
  if !seu.accountRepo.HasStock(stockId, accountId, quantity) {
    return errors.New("Insufficient Stock")
  }

  err = seu.orderExecuter.SellStock(stock,account,quantity)
  if err != nil {
    return err
  }
  // sell stock with order execution
  order := models.Order{
    StockId:     stockId,
    AccountId:   accountId,
    Quantity:     quantity,
    Price:        stock.StockPrice,
    OrderId:      time.Now().String(),
    OrderType:    models.SELL,
  }

  seu.accountRepo.AddOrder(accountId, order)

  
  return nil
}

