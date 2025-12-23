package models

import (
	"fmt"
)

type Cell struct {
	Color string
	Piece Piece
}

type ChessBoard struct {
	Board  [8][8]Cell
	Players []Player
}

// NewChessBoard creates a new chess board with the initial setup of pieces.
func NewChessBoard(playerA Player, playerB Player) *ChessBoard {
	board := &ChessBoard{
		Players: []Player{playerA,playerB},
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				board.Board[i][j].Color = "White"
			} else {
				board.Board[i][j].Color = "Black"
			}
		}
	}
	// Set up initial pieces
	board.Board[0][0].Piece = Rook{Color: "Black"}
	board.Board[0][1].Piece = Knight{Color: "Black"}
	board.Board[0][2].Piece = Bishop{Color: "Black"}
	board.Board[0][3].Piece = Queen{Color: "Black"}
	board.Board[0][4].Piece = King{Color: "Black"}
	board.Board[0][5].Piece = Bishop{Color: "Black"}
	board.Board[0][6].Piece = Knight{Color: "Black"}
	board.Board[0][7].Piece = Rook{Color: "Black"}
	board.Board[7][0].Piece = Rook{Color: "White"}
	board.Board[7][1].Piece = Knight{Color: "White"}
	board.Board[7][2].Piece = Bishop{Color: "White"}
	board.Board[7][3].Piece = Queen{Color: "White"}
	board.Board[7][4].Piece = King{Color: "White"}
	board.Board[7][5].Piece = Bishop{Color: "White"}
	board.Board[7][6].Piece = Knight{Color: "White"}
	board.Board[7][7].Piece = Rook{Color: "White"}
	for i := 0; i < 8; i++ {
		board.Board[1][i].Piece = Pawn{Color:"Black"}
		board.Board[6][i].Piece = Pawn{Color:"White"}
	}
	return board
}

// PrintBoard prints the chess board to the console.
func (b *ChessBoard) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cell := b.Board[i][j]
				if cell.Piece == nil {
					fmt.Print("  ")
					continue
				}
				switch cell.Piece.GetType() {
				case "Pawn":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♙ ")
					} else {
						fmt.Print("♟ ")
					}
				case "Knight":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♘ ")
					} else {
						fmt.Print("♞ ")
					}
				case "Bishop":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♗ ")
					} else {
						fmt.Print("♝ ")
					}
				case "Rook":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♖ ")
					} else {
						fmt.Print("♜ ")
					}
				case "Queen":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♕ ")
					} else {
						fmt.Print("♛ ")
					}
				case "King":
					if cell.Piece.GetType() == "White" {
						fmt.Print("♔ ")
					} else {
						fmt.Print("♚ ")
					}
				default:
					fmt.Print("  ")
				}
			}
		  fmt.Println()
		}
		fmt.Println()
}


func (b *ChessBoard)StartGame() {
	
	// Start game logic
	currentPlayerIndex := 0
	for {
		currentPlayer := b.Players[currentPlayerIndex]
		fmt.Printf("%s's turn\n", currentPlayer.Name)
		b.PrintBoard()

		// Get move from player
		var startRow, startCol, endRow, endCol int
		fmt.Print("Enter start row and column (e.g., 1 2): ")
		fmt.Scanln(&startRow, &startCol)
		fmt.Print("Enter end row and column (e.g., 3 4): ")
		fmt.Scanln(&endRow, &endCol)
	
		// Check if move is valid
		if !b.IsValidMove(startRow, startCol, endRow, endCol) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		// Make the move
		b.MakeMove(startRow, startCol, endRow, endCol)

		// Check if game is over
		if b.IsGameOver() {
			fmt.Println("Game Over!")
			if b.IsCheckmate() {
				fmt.Printf("%s wins by Checkmate!\n", b.Players[1-currentPlayerIndex].Name)
			} else {
				fmt.Println("Stalemate!")
			}
			break
		}

		// Switch to next player
		currentPlayerIndex = (currentPlayerIndex + 1) % 2
	}
}

// IsGameOver checks if the game is over (checkmate, stalemate, or draw).
func (b *ChessBoard) IsGameOver() bool {
	// Check for checkmate or stalemate
	// TODO: Implement checkmate and stalemate logic
	return false
}

// IsCheckmate checks if the current player is in checkmate.
func (b *ChessBoard) IsCheckmate() bool {
	// TODO: Implement checkmate logic
	return false
}

// MakeMove moves a piece from one cell to another.
func (b *ChessBoard) MakeMove(startRow, startCol, endRow, endCol int) {
	// Get the piece to move
	piece := b.Board[startRow][startCol].Piece

	// Move the piece to the new cell
	b.Board[endRow][endCol].Piece = piece
	b.Board[startRow][startCol].Piece = nil
}

// IsValidMove checks if a move is valid according to the rules of chess.
func (b *ChessBoard) IsValidMove(startRow, startCol, endRow, endCol int) bool {
	// Check if the start and end cells are within the board
	if startRow < 0 || startRow > 7 || startCol < 0 || startCol > 7 || endRow < 0 || endRow > 7 || endCol < 0 || endCol > 7 {
		return false
	}

	// Get the piece to move
	piece := b.Board[startRow][startCol].Piece

	// Check if there is a piece to move
	if piece == nil {
		return false
	}
	return true
}
