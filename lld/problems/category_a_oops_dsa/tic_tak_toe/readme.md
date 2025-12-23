Requirements:
1. The Tic-Tac-Toe game should be played on a 3x3 grid.
2. Two players take turns marking their symbols (X or O) on the grid.
3. The first player to get three of their symbols in a row (horizontally, vertically, or diagonally) wins the game.
4. If all the cells on the grid are filled and no player has won, the game ends in a draw.
5. The game should have a user interface to display the grid and allow
players to make their moves. -> wtf ??
6. The game should handle player turns and validate moves to ensure they are legal.
  - only on not filled box, correct position.
7. The game should detect and announce the winner or a draw at the end of the game.
  - what if one is already won ? before end -> its actually end :)




 Rough:

 - models
   - tictactoe
     - matrix 3*3.
     - value -> 'X'|'O'
   - Player
     - Name
     - CoinType = 'X'/'O'
 - usecase
   - tictoctoeGame
     - tictectoe
     - two players -> player one and player two
   - Register two people
   - Play game with, starting as player one
   - Print matrix
   - Validate move.
   - Is this won position
   - end game.