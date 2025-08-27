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
			game.Init(tt.input.board)
			assert.Equal(t, tt.want, game.board)
		})
	}
}
