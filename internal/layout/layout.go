package layout

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/game"
)

const (
	gridSize     = 32
	Rows         = 10
	Cols         = 10
	PanelHeight  = 36 // 上方面板高度
	PaddingX     = 32 // 面板內文字左邊距
	PaddingY     = 20 // 面板
	ScreenHeight = PanelHeight + gridSize*Rows
	ScreenWidth  = gridSize * Cols
	MineCounts   = 9
)

var buttonRectRelativePos = image.Rect(0, 0, 32, 32) // 一個方格大小的　button

type Coord struct {
	Row int
	Col int
}
type GameLayout struct {
	gameInstance *game.Game
	ClickCoord   *Coord
}

func NewGameLayout(gameInstance *game.Game) *GameLayout {
	return &GameLayout{gameInstance: gameInstance, ClickCoord: &Coord{}}
}

func (g *GameLayout) Update() error {
	// 偵測　restart icon 有被點擊
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		if xPos >= ((ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X) &&
			xPos <= (ScreenWidth)/2+buttonRectRelativePos.Max.X+0.5*gridSize &&
			yPos >= buttonRectRelativePos.Min.Y &&
			yPos <= buttonRectRelativePos.Max.Y+3 {
			g.Restart()
		}
	}
	// 當狀態為遊戲結束
	if g.gameInstance.IsGameOver || g.gameInstance.IsPlayerWin {
		return nil
	}
	// 偵測 mouse 左鍵 click 事件
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.handlePositionClickEvent(func(row, col int) {
			// 檢查是否踩到地雷
			if g.gameInstance.Board.GetCell(row, col).IsMine {
				g.gameInstance.IsGameOver = true
			}
			// 執行 Flood Fill - 更新踩到之後的更新
			g.gameInstance.Board.Reveal(row, col)
			// 檢查是否達到勝利條件
			if !g.gameInstance.IsGameOver {
				g.gameInstance.IsPlayerWin = g.gameInstance.Board.CheckIsPlayerWin()
			}
		})
	}
	// 偵測 mouse 右鍵 click 事件
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		// 標記該位置格子
		g.handlePositionClickEvent(func(row, col int) {
			g.gameInstance.Board.ToggleFlag(row, col)
		})
	}
	return nil
}

// drawUnRevealedCell - 畫出沒有被掀開的格子
func (g *GameLayout) drawUnRevealedCell(screen *ebiten.Image, row, col int) {
	vector.DrawFilledRect(
		screen,
		float32(col*gridSize),
		float32(PanelHeight+row*gridSize),
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
		float32(PanelHeight+row*gridSize),
		gridSize-1,
		gridSize-1,
		color.RGBA{200, 200, 200, 0xff},
		false,
	)
}

// drawRevealMineBackground - 畫出 click 之後 Mine 背景
func (g *GameLayout) drawRevealMineBackground(screen *ebiten.Image, row, col int) {
	bgColor := color.RGBA{200, 200, 0, 0xff}
	if g.ClickCoord.Row == row && g.ClickCoord.Col == col {
		bgColor = color.RGBA{200, 0, 0, 0xff}
	}
	vector.DrawFilledRect(
		screen,
		float32(col*gridSize),
		float32(PanelHeight+row*gridSize),
		gridSize-1,
		gridSize-1,
		bgColor,
		false,
	)
}

// drawTouchCellAdjacency - 畫出 click 之後顯示出來的值
func (g *GameLayout) drawTouchCellAdjacency(screen *ebiten.Image, row, col, value int) {
	// 繪製數字 (置中)
	textValue := fmt.Sprintf("%d", value)
	textXPos := col*gridSize + (gridSize)/2
	textYPos := PanelHeight + row*gridSize + (gridSize)/2
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
	textValue := "💣"
	textXPos := col*gridSize + (gridSize)/2
	textYPos := PanelHeight + row*gridSize + (gridSize)/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(IsMine))
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, textOpts)
}

// drawFlag - 標示 flag
func (g *GameLayout) drawFlag(screen *ebiten.Image, row, col int) {
	// 繪製數字 (置中)
	textValue := "🚩"
	textXPos := col*gridSize + (gridSize)/2
	textYPos := PanelHeight + row*gridSize + (gridSize)/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(IsFlag))
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, textOpts)
}

// drawUnRevealLogic - 繪製沒有掀開 cell 邏輯
func (g *GameLayout) drawUnRevealLogic(screen *ebiten.Image, row, col int, cell *game.Cell) {
	g.drawUnRevealedCell(screen, row, col)
	if cell.Flagged {
		g.drawFlag(screen, row, col)
	}
}

// drawRevealedMineLogic - 畫出被 clicked 到地雷時的邏輯
func (g *GameLayout) drawRevealedMineLogic(screen *ebiten.Image, row, col int, cell *game.Cell) {
	g.drawRevealMineBackground(screen, row, col)
	if !cell.Flagged {
		g.drawTouchCellMine(screen, row, col)
	} else {
		g.drawFlag(screen, row, col)
	}
}

// drawRevealedCell - 畫出被 clicked 到的格子
func (g *GameLayout) drawRevealedCell(screen *ebiten.Image, row, col int, cell *game.Cell) {
	g.drawTouchCellBackground(screen, row, col)
	if cell.AdjacenetMines != 0 {
		g.drawTouchCellAdjacency(screen, row, col, cell.AdjacenetMines)
	}
	if cell.IsMine {
		g.drawRevealedMineLogic(screen, row, col, cell)
	}
}

// drawBoard - 畫出目前盤面狀態
func (g *GameLayout) drawBoard(screen *ebiten.Image) {
	for row := 0; row < Rows; row++ {
		for col := 0; col < Cols; col++ {
			// 取出格子狀態
			cell := g.gameInstance.Board.GetCell(row, col)

			// 根據格子狀態，顯示對應的畫面
			// 當格子沒有被掀開時,畫出原本的灰階
			if !cell.Revealed {
				g.drawUnRevealLogic(screen, row, col, cell)
			} else {
				g.drawRevealedCell(screen, row, col, cell)
			}
		}
	}
}

// drawGamePanel - 繪製遊戲狀態面板
func (g *GameLayout) drawGamePanel(screen *ebiten.Image) {
	emojiIcon, bgColor := g.getColorStatus()
	panel := ebiten.NewImage(ScreenWidth, PanelHeight)
	panel.Fill(bgColor)
	screen.DrawImage(panel, nil)
	// 畫旗子面板（固定在左方）
	g.drawRemainingFlagInfo(screen)
	// 畫顯示狀態　button
	g.drawButtonWithIcon(screen, emojiIcon)
}

// drawButtonWithIcon　- 繪製 buttonIcon
func (g *GameLayout) drawButtonWithIcon(screen *ebiten.Image, emojiIcon string) {
	vector.DrawFilledRect(screen,
		float32((ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X),
		float32(buttonRectRelativePos.Min.Y),
		float32(buttonRectRelativePos.Dx()+0.5*gridSize),
		float32(buttonRectRelativePos.Dy()+3),
		color.RGBA{120, 120, 120, 255},
		true,
	)
	vector.DrawFilledCircle(screen, ScreenWidth/2, gridSize/2, 16,
		color.RGBA{180, 180, 0, 255},
		true,
	)
	emojiValue := emojiIcon
	emojiXPos := (ScreenWidth) / 2
	emojiYPos := PaddingY
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getTileColor(IsButtonIcon))
	emojiOpts.PrimaryAlign = text.AlignCenter
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   32,
	}, emojiOpts)
}

// drawRemainingFlagInfo
func (g *GameLayout) drawRemainingFlagInfo(screen *ebiten.Image) {
	// 畫旗子面板（固定在左方）
	textValue := fmt.Sprintf("%03d", g.gameInstance.Board.GetRemainingFlags())
	textXPos := PaddingX + len(textValue)
	textYPos := PaddingY
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(-1))
	textOpts.PrimaryAlign = text.AlignStart
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   20,
	}, textOpts)
	emojiValue := "🚩"
	emojiXPos := len(emojiValue)
	emojiYPos := PaddingY
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getTileColor(-1))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   20,
	}, emojiOpts)
}

func (g *GameLayout) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
	g.drawGamePanel(screen)
}

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// getColorStatus - 根據 IsGameOver 與 IsPlayerWin 來找出對 message, bgColor
func (g *GameLayout) getColorStatus() (string, color.RGBA) {
	bgColor := color.RGBA{100, 100, 0x10, 0xFF}
	status := "😀"
	if g.gameInstance.IsGameOver {
		status = "😵"
		bgColor = color.RGBA{150, 0, 0x10, 0xFF}
	}
	if g.gameInstance.IsPlayerWin {
		status = "😎"
		bgColor = color.RGBA{200, 200, 0, 0xFF}
	}
	return status, bgColor
}

// handlePositionClickEvent - 處理 click 之後把 positon 傳入
func (g *GameLayout) handlePositionClickEvent(listenHandler func(row, col int)) {
	xPos, yPos := ebiten.CursorPosition()
	// 當在面板下方才處理
	if yPos >= PanelHeight {
		row := (yPos - PanelHeight) / gridSize
		col := xPos / gridSize
		g.ClickCoord.Row = row
		g.ClickCoord.Col = col
		if row >= 0 && row < Rows && col >= 0 && col < Cols {
			listenHandler(row, col)
		}
	}
}

// Restart - 重新建立 Game 狀態
func (g *GameLayout) Restart() {
	g.gameInstance = game.NewGame(Rows, Cols, MineCounts)
}
