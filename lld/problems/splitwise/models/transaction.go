package models

type Transaction struct{
  From string   `json:"from"`
  To string     `json:"to"`
  Amount float64  `json:"Amount"`
}