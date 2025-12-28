package main

import (
	"main/models"
)

func main() {
	chess := models.NewChessBoard(models.Player{ID: 1 ,Name:"anubhav" }, models.Player{ID: 2, Name: "Gupta"})
	chess.PrintBoard()

	chess.StartGame()
}
