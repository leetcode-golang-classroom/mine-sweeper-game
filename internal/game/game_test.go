package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameInit(t *testing.T) {
	type field struct {
		board       *Board
		rows        int
		cols        int
		minesNumber int
	}
	tests := []struct {
		name  string
		input field
		want  *Board
	}{
		{
			name: "Empty Board with row = 5, col = 5, minesNumer = 5",
			input: field{
				rows:        5,
				cols:        5,
				minesNumber: 5,
				board: &Board{
					rows: 5,
					cols: 5,
					cells: [][]*Cell{
						{
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         true,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
						},
						{
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         true,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
						},
						{
							{
								IsMine:         true,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
						},
						{
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         true,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
						},
						{
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         true,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 0,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 2,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
							{
								IsMine:         false,
								Revealed:       false,
								Flagged:        false,
								AdjacenetMines: 1,
							},
						},
					},
				},
			},
			want: &Board{
				rows: 5,
				cols: 5,
				cells: [][]*Cell{
					{
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         true,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
					},
					{
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         true,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
					},
					{
						{
							IsMine:         true,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
					},
					{
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         true,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
					},
					{
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         true,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							Revealed:       false,
							Flagged:        false,
							AdjacenetMines: 1,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.input.rows, tt.input.cols, tt.input.minesNumber)
			game.Init(tt.input.board, func(coords []coord) {})
			assert.Equal(t, tt.want.cells, game.Board.cells)
		})
	}
}

func TestCalculateAdjacentMines(t *testing.T) {
	type field struct {
		board       *Board
		rows        int
		cols        int
		minesNumber int
	}
	tests := []struct {
		name  string
		input field
		want  *Board
	}{
		{
			name: "Test CalculateAdjacentMines with specific input board, row = 5, col = 5, mineNumber = 4",
			input: field{
				rows:        5,
				cols:        5,
				minesNumber: 4,
				board: &Board{
					rows:                 5,
					cols:                 5,
					minePositionShuffler: func(coords []coord) {},
					cells: [][]*Cell{
						{
							{
								IsMine: false,
							},
							{
								IsMine: true,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
						},
						{
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: true,
							},
							{
								IsMine: false,
							},
						},
						{
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: true,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
						},
						{
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
						},
						{
							{
								IsMine: true,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
							{
								IsMine: false,
							},
						},
					},
				},
			},
			want: &Board{
				rows:                 5,
				cols:                 5,
				minePositionShuffler: func(coords []coord) {},
				cells: [][]*Cell{
					{
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         true,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
					},
					{
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							AdjacenetMines: 3,
						},
						{
							IsMine:         true,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
					},
					{
						{
							IsMine:         false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         true,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
					},
					{
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 2,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 0,
						},
					},
					{
						{
							IsMine:         true,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 1,
						},
						{
							IsMine:         false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 0,
						},
						{
							IsMine:         false,
							AdjacenetMines: 0,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.input.rows, tt.input.cols, tt.input.minesNumber)
			game.Init(tt.input.board, func(coords []coord) {})
			game.Board.CalculateAdjacentMines()
			assert.Equal(t, tt.want.cells, game.Board.cells)
		})
	}
}
