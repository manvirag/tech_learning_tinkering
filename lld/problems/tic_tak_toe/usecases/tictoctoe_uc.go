package usecases

import (
	"fmt"
	"main/models"
)

type TicTocToeGameInterface interface {
	Play() error
}

type TicTocToeGame struct {
	TicTocToeBoard *models.TicTacToe
	Players        []models.Player
}

func NewTicTocToeGame(playerA models.Player, playerB models.Player) *TicTocToeGame {
	return &TicTocToeGame{
		TicTocToeBoard: models.NewTicTacToe(),
		Players:        []models.Player{ playerA, playerB},
	}
}

func (ttt *TicTocToeGame) Play() error {

	for {

    for i:=0;i<2;i++ {
        ttt.TicTocToeBoard.PrintBoard()
        fmt.Printf("Player %s, please enter your move: ", ttt.Players[i].Name)
        var row, col int
        fmt.Scanf("%d", &row)
        fmt.Scanf("%d", &col)
      // fmt.Println(row)
      // fmt.Println(col)
        err := ttt.TicTocToeBoard.ValidateMove(row, col, ttt.Players[i].CoinType)
        if err != nil {
          fmt.Printf("Please enter valid move. Player %s", ttt.Players[i].Name)
          i = i-1;
          continue
        }
        ttt.TicTocToeBoard.MakeMove(row, col, ttt.Players[i].CoinType)
        isGameOver, draw := ttt.TicTocToeBoard.IsGameOver(row, col)
        if isGameOver {
          fmt.Println("Thanks for playing. \nGame is Over Player\n.")
          if draw {
            fmt.Println("Game is Draw. Please Play Again \n")
          } else {
            fmt.Printf("Player %s won the game.\n", ttt.Players[i].Name)
          }
          break
        }
    }
		
	}
	return nil
}
