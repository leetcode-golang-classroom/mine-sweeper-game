package game

// Cell - 單一格子
type Cell struct {
	IsMine         bool // 是否為地雷
	Revealed       bool // 是否被翻開
	Flagged        bool // 是否被插旗
	AdjacenetMines int  // 周圍地雷數
}

// Board - 棋盤
type Board struct {
	rows                 int              // 總共格數
	cols                 int              // 總共列數
	cells                [][]*Cell        // 整格棋盤狀態
	minePositionShuffler positionShuffler // 亂序器用來安排地雷格子
}

// Game - 遊戲物件
type Game struct {
	Board       *Board // 棋盤物件
	isGameOver  bool   // 是否遊戲結束
	isPlayerWin bool   // 玩家是否獲勝
	remaining   int    // 剩餘需要翻開的格子數
}

// coord - 紀錄該格字座標
type coord struct {
	Row int
	Col int
}

// positionShuffler - 亂序器用來安排地雷格子
type positionShuffler func(coords []coord)

func NewGame(rows, cols, mineCount int) *Game {
	board := NewBoard(rows, cols)
	// 設定地雷
	board.PlaceMines(mineCount)
	// 計算結果
	board.CalculateAdjacentMines()
	return &Game{
		Board:       board,
		remaining:   rows*cols - mineCount,
		isGameOver:  false,
		isPlayerWin: false,
	}
}

// NewBoard - 初始化盤面
func NewBoard(rows, cols int) *Board {
	board := &Board{
		rows:                 rows,
		cols:                 cols,
		minePositionShuffler: defaultPositionShuffler,
	}
	board.cells = make([][]*Cell, rows)
	for row := range board.cells {
		board.cells[row] = make([]*Cell, cols)
		for col := range board.cells[row] {
			board.cells[row][col] = &Cell{}
		}
	}
	return board
}

func (g *Game) Init(board *Board, minePositionShuffler positionShuffler) {
	if minePositionShuffler != nil {
		g.Board.minePositionShuffler = minePositionShuffler
	}
	// 無效的設定
	if board == nil || len(board.cells) != board.rows || len(board.cells[0]) != board.cols {
		return
	}
	// 設定資料
	for row := range board.cells {
		for col := range board.cells[row] {
			sourceCell := board.cells[row][col]
			g.Board.cells[row][col].AdjacenetMines = sourceCell.AdjacenetMines
			g.Board.cells[row][col].IsMine = sourceCell.IsMine
			g.Board.cells[row][col].Revealed = sourceCell.Revealed
			g.Board.cells[row][col].Flagged = sourceCell.Flagged
		}
	}
}

// PlaceMines - 使用 minePositionShuffler 選出 mineCount 個地雷
func (b *Board) PlaceMines(mineCount int) {
	if mineCount < 0 {
		return
	}
	// 蒐集所有 coord
	coords := make([]coord, 0, b.cols*b.rows)
	for row := range b.cells {
		for col := range b.cells[row] {
			coords = append(coords, coord{Row: row, Col: col})
		}
	}
	// 使用 minePositionShuffler 作洗牌
	b.minePositionShuffler(coords)
	coordLen := len(coords)
	// 避免 mineCount 超過 coords 個數
	if mineCount > coordLen {
		mineCount = coordLen
	}
	// 設定前 mineCount 為地雷
	for i := 0; i < mineCount; i++ {
		b.cells[coords[i].Row][coords[i].Col].IsMine = true
	}
}

// CalculateAdjacentMines - 計算鄰近地雷個數
func (b *Board) CalculateAdjacentMines() {
	// 鄰近所有方向
	neighborDirections := [8]coord{
		{Row: -1, Col: -1}, {Row: -1, Col: 0}, {Row: -1, Col: 1},
		{Row: 0, Col: -1}, {Row: 0, Col: 1},
		{Row: 1, Col: -1}, {Row: 1, Col: 0}, {Row: 1, Col: 1},
	}
	for row := range b.cells {
		for col := range b.cells[row] {
			// 當遇到地雷格時 跳過
			if b.cells[row][col].IsMine {
				continue
			}
			// 開始累計鄰近的地雷數
			accumCount := 0
			for _, direction := range neighborDirections {
				neighborRow, neighborCol := row+direction.Row, col+direction.Col
				if neighborRow >= 0 && neighborRow < b.rows &&
					neighborCol >= 0 && neighborCol < b.cols &&
					b.cells[neighborRow][neighborCol].IsMine {
					accumCount++
				}
			}
			b.cells[row][col].AdjacenetMines = accumCount
		}
	}
}

func (board *Board) GetCell(row, col int) *Cell {
	return board.cells[row][col]
}
