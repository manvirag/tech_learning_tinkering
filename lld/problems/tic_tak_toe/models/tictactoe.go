package models

import ("fmt"
        "errors"
       )

type CoinType string

const (
  X CoinType = "X"
  O CoinType = "O"
  E CoinType = "E"
)

type TicTacToeInterface interface{
  PrintBoard()
  ValidateMove(row int, col int, coinType CoinType) error
  MakeMove(row int, col int, coinType CoinType)
  IsGameOver( row int , col int) (bool, bool)
  CheckLeftDiagonal () bool
  CheckRightDiagonal () bool
}

type TicTacToe struct {
  Board [3][3]CoinType
}

var invalidDiagonalPos = [][]int{{1, 0}, {0, 1}, {1, 2}, {2, 1}}

func NewTicTacToe() *TicTacToe{

  var board [3][3]CoinType

  for i:= 0 ;i<3;i++{
    for j:=0;j<3;j++{
      board[i][j] = E
    }
  }
  
  return &TicTacToe{
    Board: board,
  }
}


func (t *TicTacToe) PrintBoard(){
      for i:= 0 ;i<3;i++{
        for j:=0;j<3;j++{
      fmt.Printf("%s ", t.Board[i][j])
    }
    fmt.Println()
  }
}

func (t *TicTacToe) ValidateMove(row int, col int, coinType CoinType) error {
    if row >= 3 || col >= 3 {
      return errors.New("Row or Column is out of bounds")
    }
    coinValue := t.Board[row][col]
    if coinValue != E{
      return errors.New("This place is not empty please enter valid move")
    }
    t.Board[row][col] = coinType
    return nil  
}

func (t * TicTacToe) MakeMove(row int, col int, coinType CoinType) {
    t.Board[row][col] = coinType
}

func (t * TicTacToe) IsGameOver(row int, col int) (bool, bool){
   horizontalMatch := true
   // Check Horizontal
    for i:=0;i<3;i++{
       if(t.Board[row][i] != t.Board[row][col]){
          horizontalMatch = false
       }
    }
  verticalMatch := true
   // Check Vertical
    for i:=0;i<3;i++{
       if(t.Board[i][col] != t.Board[row][col]){
          verticalMatch = false
       }
    }

  diagonalMatch := true
    // Check Diagonal
   fmt.Println(row)
  fmt.Println(col)
   for i:=0;i<4;i++{
        if row == invalidDiagonalPos[i][0] && col == invalidDiagonalPos[i][1] {
          diagonalMatch = false
          break;
        }
    
   }
   fmt.Println(diagonalMatch)
   if diagonalMatch == true {
        if row == 1 && col == 1 {
          checkLeftDiagonal := t.CheckLeftDiagonal();
          checkRightDiagonal := t.CheckRightDiagonal();
          diagonalMatch = checkLeftDiagonal || checkRightDiagonal
        } else if (row == 0 && col == 0) || (row == 2 && col == 2) {
          diagonalMatch = t.CheckLeftDiagonal()
        }  else {
           diagonalMatch  = t.CheckRightDiagonal()
        }

   }
    fmt.Println(horizontalMatch)
    fmt.Println(verticalMatch)
    fmt.Println(diagonalMatch)
    if horizontalMatch || verticalMatch || diagonalMatch {
      return true, false
    }

    for i:=0;i<3;i++{
      for j:=0;j<3;j++{
        if t.Board[i][j] == E{
            return false, false
        }
      }
    }

   return true, true
}

func (t *TicTacToe)CheckLeftDiagonal() bool{
  return t.Board[0][0] == t.Board[1][1] && t.Board[1][1] == t.Board[2][2] && t.Board[0][0] != E;
}


func (t *TicTacToe)CheckRightDiagonal() bool{
  return t.Board[0][2] == t.Board[1][1] && t.Board[1][1] == t.Board[2][0] && t.Board[0][0] != E;
}




