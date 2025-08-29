package layout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/game"
)

const (
	gridSize     = 32
	Rows         = 10
	Cols         = 10
	ScreenWidth  = gridSize * Rows
	ScreenHeight = gridSize * Cols
	MineCounts   = 9
)

type GameLayout struct {
	gameInstance *game.Game
}

func NewGameLayout(gameInstance *game.Game) *GameLayout {
	return &GameLayout{gameInstance: gameInstance}
}
func (g *GameLayout) Update() error {
	return nil
}

func (g *GameLayout) drawUnTouchCell(screen *ebiten.Image, row, col int) {
	vector.DrawFilledRect(
		screen,
		float32(row*gridSize),
		float32(col*gridSize),
		gridSize-1,
		gridSize-1,
		color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
		false,
	)
}
func (g *GameLayout) Draw(screen *ebiten.Image) {
	for row := 0; row < Rows; row++ {
		for col := 0; col < Cols; col++ {
			cell := g.gameInstance.Board.GetCell(row, col)
			if !cell.Revealed {
				g.drawUnTouchCell(screen, row, col)
			}
		}
	}
}

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
