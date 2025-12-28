package models

import "errors"

type StockExchangeRepository struct {
	Stocks []Stock
}

type StockExchangeRepositoryInterface interface {
	GetStocks() []Stock
	AddStock(stock Stock)
	UpadteStock(stock Stock)
  GetStock(stockId string) (Stock, error)
}

func NewStockExchangeRepository() *StockExchangeRepository {
	return &StockExchangeRepository{
		Stocks: make([]Stock, 0),
	}
}

func (m *StockExchangeRepository) GetStocks() []Stock {
	return m.Stocks
}

func (m *StockExchangeRepository) AddStock(stock Stock) {
	m.Stocks = append(m.Stocks, stock)
}

func (m * StockExchangeRepository) GetStock(stockId string) (Stock, error){
  for _, stock := range m.Stocks {
    if (stock.StockId == stockId) {
      return stock, nil
    }
  }
  return Stock{}, errors.New("Stock not found")
}

func (m *StockExchangeRepository) UpadteStock(stock Stock) {
	for i, s := range m.Stocks {
		if s.StockId == stock.StockId {
			m.Stocks[i] = stock
			return
		}
	}
}
