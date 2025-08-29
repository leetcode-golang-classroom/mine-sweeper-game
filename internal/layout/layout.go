package layout

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	// 偵測 mouse click 事件
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		row := yPos / gridSize
		col := xPos / gridSize
		if row >= 0 && row < Rows && col >= 0 && col < Cols {
			// 執行 Flood Fill
			g.gameInstance.Board.Reveal(row, col)
		}
	}
	return nil
}

// drawUnTouchCell - 畫出沒有被掀開的格子
func (g *GameLayout) drawUnTouchCell(screen *ebiten.Image, row, col int) {
	vector.DrawFilledRect(
		screen,
		float32(col*gridSize),
		float32(row*gridSize),
		gridSize-1,
		gridSize-1,
		color.RGBA{100, 100, 100, 0xff},
		false,
	)
}

// drawTouchCellBackground - 畫出 click 之後背景
func (g *GameLayout) drawTouchCellBackground(screen *ebiten.Image, row, col int) {
	vector.DrawFilledRect(
		screen,
		float32(col*gridSize),
		float32(row*gridSize),
		gridSize-1,
		gridSize-1,
		color.RGBA{200, 200, 200, 0xff},
		false,
	)
}

// drawTouchCellAdjacency - 畫出 click 之後顯示出來的值
func (g *GameLayout) drawTouchCellAdjacency(screen *ebiten.Image, row, col, value int) {
	// 繪製數字 (置中)
	textValue := fmt.Sprintf("%d", value)
	textXPos := col*gridSize + (gridSize)/2
	textYPos := row*gridSize + (gridSize)/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(value))
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

// drawTouchCellMine - 畫出地雷
func (g *GameLayout) drawTouchCellMine(screen *ebiten.Image, row, col int) {
	// 繪製數字 (置中)
	textValue := "X"
	textXPos := col*gridSize + (gridSize)/2
	textYPos := row*gridSize + (gridSize)/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(-1))
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

func (g *GameLayout) Draw(screen *ebiten.Image) {
	for row := 0; row < Rows; row++ {
		for col := 0; col < Cols; col++ {
			// 取出格子狀態
			cell := g.gameInstance.Board.GetCell(row, col)

			// 根據格子狀態，顯示對應的畫面
			// 當格子沒有被掀開時,畫出原本的灰階
			if !cell.Revealed {
				g.drawUnTouchCell(screen, row, col)
			} else {
				g.drawTouchCellBackground(screen, row, col)
				if cell.AdjacenetMines != 0 {
					g.drawTouchCellAdjacency(screen, row, col, cell.AdjacenetMines)
				}
				if cell.IsMine {
					g.drawTouchCellMine(screen, row, col)
				}
			}
		}
	}
}

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
