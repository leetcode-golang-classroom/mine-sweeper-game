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
	gridSize    = 32
	PanelHeight = 72 // 上方面板高度
	PaddingX    = 32 // 面板內文字左邊距
	PaddingY    = 20 // 面板
)

var DefaultRows = LevelSetupMap[Easy].Rows
var DefaultCols = LevelSetupMap[Easy].Cols
var DefaultScreenHeight = PanelHeight + gridSize*DefaultRows
var DefaultScreenWidth = gridSize * DefaultCols
var DefaultMineCounts = LevelSetupMap[Easy].MineCounts
var buttonRectRelativePos = image.Rect(0, 0, 32, 32) // 一個方格大小的　button

type Coord struct {
	Row int
	Col int
}

// 遊戲畫面狀態
type GameLayout struct {
	gameInstance *game.Game //　遊戲物件
	ClickCoord   *Coord     //　使用者點擊座標
	elapsedTime  int        // 經過時間
	Rows         int        // 紀錄遊戲 Row 大小
	Cols         int        // 紀錄遊戲 Col 大小
	MineCounts   int        // 紀錄 MineCounts
	ScreenHeight int
	ScreenWidth  int
	level        Level
}

func NewGameLayout(gameInstance *game.Game) *GameLayout {
	return &GameLayout{gameInstance: gameInstance, ClickCoord: &Coord{},
		Rows:       gameInstance.Board.Rows,
		Cols:       gameInstance.Board.Cols,
		MineCounts: gameInstance.MineCounts,
		level:      Easy,
	}
}

func (g *GameLayout) Update() error {
	// 偵測　level icon 有被點擊
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		if xPos >= ((g.ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X) &&
			xPos <= (g.ScreenWidth)/2+buttonRectRelativePos.Max.X+0.5*gridSize &&
			yPos >= buttonRectRelativePos.Min.Y &&
			yPos <= buttonRectRelativePos.Max.Y+3 {
			g.ChangeLevel()
			g.Restart()
		}
	}
	// 偵測　restart icon 有被點擊
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		if xPos >= ((g.ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X) &&
			xPos <= (g.ScreenWidth)/2+buttonRectRelativePos.Max.X+0.5*gridSize &&
			yPos >= gridSize+buttonRectRelativePos.Min.Y &&
			yPos <= gridSize+buttonRectRelativePos.Max.Y+3 {
			g.Restart()
		}
	}
	// 當遊戲還沒停止時，就更新經過時間
	if !g.gameInstance.IsGameOver && !g.gameInstance.IsPlayerWin {
		g.elapsedTime = g.gameInstance.GetElapsedTime()
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
	textOpts.ColorScale.ScaleWithColor(getTileColor(-1))
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
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
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
	panel := ebiten.NewImage(g.ScreenWidth, PanelHeight)
	panel.Fill(bgColor)
	screen.DrawImage(panel, nil)
	// 畫旗子面板（固定在左方）
	g.drawRemainingFlagInfo(screen)
	// 畫顯示狀態　button
	g.drawButtonWithIcon(screen, emojiIcon)
	// 畫出經過時間
	g.drawElaspedTimeInfo(screen)
	// 畫出 Level Info Button
	g.drawLevelInfo(screen)
}

func (g *GameLayout) drawLevelInfo(screen *ebiten.Image) {
	// 畫Level（固定在左方）
	textValue := "Level:"
	textXPos := len(textValue)
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
	emojiIcon := LevelIconMap[g.level]
	g.drawLevelButtonWithIcon(screen, emojiIcon)
}

func (g *GameLayout) drawLevelButtonWithIcon(screen *ebiten.Image, emojiIcon string) {
	vector.DrawFilledRect(screen,
		float32((g.ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X),
		float32(buttonRectRelativePos.Min.Y),
		float32(buttonRectRelativePos.Dx()+0.5*gridSize),
		float32(buttonRectRelativePos.Dy()+4),
		color.RGBA{120, 120, 120, 255},
		true,
	)
	vector.DrawFilledCircle(screen, float32(g.ScreenWidth/2), gridSize/2, 16,
		LevelColorMap[g.level],
		true,
	)
	emojiValue := emojiIcon
	emojiXPos := (g.ScreenWidth) / 2
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

// drawButtonWithIcon　- 繪製 buttonIcon
func (g *GameLayout) drawButtonWithIcon(screen *ebiten.Image, emojiIcon string) {
	vector.DrawFilledRect(screen,
		float32((g.ScreenWidth-1.5*gridSize)/2+buttonRectRelativePos.Min.X),
		float32(gridSize+buttonRectRelativePos.Min.Y),
		float32(buttonRectRelativePos.Dx()+0.5*gridSize),
		float32(buttonRectRelativePos.Dy()+4),
		color.RGBA{120, 120, 120, 255},
		true,
	)
	vector.DrawFilledCircle(screen, float32(g.ScreenWidth/2), gridSize+gridSize/2, 16,
		color.RGBA{180, 180, 0, 255},
		true,
	)
	emojiValue := emojiIcon
	emojiXPos := (g.ScreenWidth) / 2
	emojiYPos := gridSize + PaddingY
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
	textYPos := gridSize + PaddingY
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
	emojiYPos := gridSize + PaddingY
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getTileColor(IsFlag))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
}

// drawElaspedTimeInfo
func (g *GameLayout) drawElaspedTimeInfo(screen *ebiten.Image) {
	// 畫旗子面板（固定在左方）
	textValue := fmt.Sprintf("%03d", g.elapsedTime)
	textXPos := g.ScreenWidth - gridSize/2 + len(textValue)
	textYPos := gridSize + PaddingY
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(-1))
	textOpts.PrimaryAlign = text.AlignEnd
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   20,
	}, textOpts)
	emojiValue := "⏰"
	emojiXPos := g.ScreenWidth - 3*gridSize + len(emojiValue)
	emojiYPos := gridSize + PaddingY
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getTileColor(IsClock))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
}

func (g *GameLayout) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
	g.drawGamePanel(screen)
}

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.ScreenHeight = PanelHeight + gridSize*g.Rows
	g.ScreenWidth = gridSize * g.Cols
	return g.ScreenWidth, g.ScreenHeight
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
		if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
			listenHandler(row, col)
		}
	}
}

// Restart - 重新建立 Game 狀態
func (g *GameLayout) Restart() {
	g.Rows = LevelSetupMap[g.level].Rows
	g.Cols = LevelSetupMap[g.level].Cols
	g.MineCounts = LevelSetupMap[g.level].MineCounts
	g.ScreenHeight = PanelHeight + gridSize*g.Rows
	g.ScreenWidth = gridSize * g.Cols
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(fmt.Sprintf("%s Mine Sweeper Grid", LevelMessage[g.level]))
	g.gameInstance = game.NewGame(g.Rows, g.Cols, g.MineCounts)
}

func (g *GameLayout) ChangeLevel() {
	g.level = (g.level + 1) % 3
}
