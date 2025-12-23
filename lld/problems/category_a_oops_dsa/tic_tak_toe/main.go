package main

import (
	"main/models"
	"main/usecases"
)

func main() {
	 // tictoctoe := models.NewTicTacToe()
	 tictoctoe_game := usecases.NewTicTocToeGame(models.Player{Name: "anubhav", CoinType: models.X}, models.Player{Name:"xxxx", CoinType: models.O})

	tictoctoe_game.Play()
}
