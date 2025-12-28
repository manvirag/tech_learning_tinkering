package models

type PortFolio struct {
  PortFolioId string
  Stocks []string
  StockCounts map[string]int
}

