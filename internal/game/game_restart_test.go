package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameRestart(t *testing.T) {
	const (
		rows      = 5
		cols      = 5
		mineCount = 5
	)

	predicableMineShuffler := func(coords []coord) {
		// not shuffler
	}
	newGameWithPredictableMines := func() *Game {
		game := NewGame(rows, cols, mineCount)
		game.Board.minePositionShuffler = predicableMineShuffler
		game.Board.cells = make([][]*Cell, rows)
		for r := range game.Board.cells {
			game.Board.cells[r] = make([]*Cell, cols)
			for c := range game.Board.cells[r] {
				game.Board.cells[r][c] = &Cell{}
			}
		}
		game.Board.mineCoords = []coord{}
		game.Board.PlaceMines(mineCount)
		game.Board.CalculateAdjacentMines()
		return game
	}
	// 1. create an initial game state for comparison
	gameInitial := newGameWithPredictableMines()

	// 2. create another game to manipulate
	gamePlaying := newGameWithPredictableMines()

	// 3. Change the state of gamePlaying
	// Let's reveal a safe cell (a cell that is not a mine)
	// With our predictable shuffler, mines are at (0,0), (0,1), (0,2), (0,3), (0,4)
	// Let's reveal cell (4,4) which is safe
	gamePlaying.Board.Reveal(4, 4)

	assert.NotEqual(t, gameInitial.Board.cells, gamePlaying.Board.cells)

	// 5 Simulate a restart by create a new game
	restartedGame := NewGame(rows, cols, mineCount)
	restartedGame.Board.minePositionShuffler = predicableMineShuffler
	restartedGame.Board.minePositionShuffler = predicableMineShuffler
	restartedGame.Board.cells = make([][]*Cell, rows)
	for r := range restartedGame.Board.cells {
		restartedGame.Board.cells[r] = make([]*Cell, cols)
		for c := range restartedGame.Board.cells[r] {
			restartedGame.Board.cells[r][c] = &Cell{}
		}
	}
	restartedGame.Board.mineCoords = []coord{}
	restartedGame.Board.PlaceMines(mineCount)
	restartedGame.Board.CalculateAdjacentMines()

	// verify that the restarted game state is identical to intial game state
	assert.Equal(t, gameInitial.Board.cells, restartedGame.Board.cells)
	assert.Equal(t, gameInitial.IsGameOver, restartedGame.IsGameOver)
	assert.Equal(t, gameInitial.Board.GetRemainingFlags(), restartedGame.Board.GetRemainingFlags())
}
