package game

import "time"

// Cell - 單一格子
type Cell struct {
	IsMine         bool // 是否為地雷
	Revealed       bool // 是否被翻開
	Flagged        bool // 是否被插旗
	AdjacenetMines int  // 周圍地雷數
}

// Board - 棋盤
type Board struct {
	Rows                     int              // 總共格數
	Cols                     int              // 總共列數
	cells                    [][]*Cell        // 整格棋盤狀態
	minePositionShuffler     positionShuffler // 亂序器用來安排地雷格子
	remainingFlags           int              // 剩餘標記數
	mineCoords               []coord          // 紀錄被設定成 mines 的座標
	remainingUnRevealedCells int              // 剩餘需要翻開的格子數
}

// Game - 遊戲物件
type Game struct {
	Board       *Board    // 棋盤物件
	IsGameOver  bool      // 是否遊戲結束
	IsPlayerWin bool      // 玩家是否獲勝
	startTime   time.Time // 遊戲開始時間
	MineCounts  int       // minecounts
}

// coord - 紀錄該格字座標
type coord struct {
	Row int
	Col int
}

// positionShuffler - 亂序器用來安排地雷格子
type positionShuffler func(coords []coord)

func NewGame(rows, cols, mineCount int) *Game {
	board := NewBoard(rows, cols, mineCount)
	// 設定地雷
	board.PlaceMines(mineCount)
	// 計算結果
	board.CalculateAdjacentMines()
	return &Game{
		Board:       board,
		IsGameOver:  false,
		IsPlayerWin: false,
		startTime:   time.Now().UTC(),
		MineCounts:  mineCount,
	}
}

// NewBoard - 初始化盤面
func NewBoard(rows, cols, mineCount int) *Board {
	board := &Board{
		Rows:                     rows,
		Cols:                     cols,
		minePositionShuffler:     defaultPositionShuffler,
		remainingFlags:           mineCount,
		remainingUnRevealedCells: rows*cols - mineCount,
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
	if board == nil || len(board.cells) != board.Rows || len(board.cells[0]) != board.Cols {
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
	coords := make([]coord, 0, b.Cols*b.Rows)
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
		b.mineCoords = append(b.mineCoords, coords[i])
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
				if neighborRow >= 0 && neighborRow < b.Rows &&
					neighborCol >= 0 && neighborCol < b.Cols &&
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

// ToggleFlag - 標記地雷
func (board *Board) ToggleFlag(row, col int) {
	// 超出邊界
	if row < 0 || row >= board.Rows ||
		col < 0 || col >= board.Cols {
		return
	}

	cell := board.cells[row][col]
	// 已經被揭開
	if cell.Revealed {
		return
	}

	count := 0
	if cell.Flagged {
		count--
	} else {
		count++
	}
	if board.remainingFlags-count < 0 {
		return
	}
	board.remainingFlags -= count
	board.cells[row][col].Flagged = !cell.Flagged
}

// Reveal - 從 row, col 開始翻開周圍不是地雷，直到遇到非零的格子
func (board *Board) Reveal(row, col int) {
	visitQueue := []coord{{
		Row: row,
		Col: col,
	}}
	// 透過 queue 來實做 BFS
	for len(visitQueue) > 0 {
		// pop up first
		cellCoord := visitQueue[0]
		visitQueue = visitQueue[1:]
		curRow, curCol := cellCoord.Row, cellCoord.Col

		// 超出邊界
		if curRow < 0 || curRow >= board.Rows ||
			curCol < 0 || curCol >= board.Cols {
			continue
		}

		cell := board.cells[curRow][curCol]
		// 已經被揭開
		if cell.Revealed {
			continue
		}

		// 標注該格已經被揭開
		board.cells[curRow][curCol].Revealed = true
		board.remainingUnRevealedCells--
		if cell.Flagged {
			board.remainingFlags++
			board.cells[curRow][curCol].Flagged = false
		}
		if cell.IsMine {
			board.revealMines()
			return
		}
		// 如果是空白格 (AdjacenetMines = 0, 且不是地雷)
		if !cell.IsMine && cell.AdjacenetMines == 0 {
			// 鄰近所有方向
			neighborDirections := [8]coord{
				{Row: -1, Col: -1}, {Row: -1, Col: 0}, {Row: -1, Col: 1},
				{Row: 0, Col: -1}, {Row: 0, Col: 1},
				{Row: 1, Col: -1}, {Row: 1, Col: 0}, {Row: 1, Col: 1},
			}
			for _, direction := range neighborDirections {
				neighborRow, neighborCol := curRow+direction.Row, curCol+direction.Col
				visitQueue = append(visitQueue, coord{
					Row: neighborRow,
					Col: neighborCol,
				})
			}
		}
	}
}

// revealMines - 顯示所有 Mines
func (board *Board) revealMines() {
	for _, mineCoord := range board.mineCoords {
		if !board.cells[mineCoord.Row][mineCoord.Col].Flagged &&
			!board.cells[mineCoord.Row][mineCoord.Col].Revealed {
			board.cells[mineCoord.Row][mineCoord.Col].Revealed = true
		}
		if board.cells[mineCoord.Row][mineCoord.Col].Flagged {
			board.cells[mineCoord.Row][mineCoord.Col].Revealed = true
		}
	}
}

// CheckIsPlayerWin - 檢查是否所有該非地雷的格子都有被掀開
func (board *Board) CheckIsPlayerWin() bool {
	return board.remainingUnRevealedCells == 0
}

// GetRemainingFlags - 取出有標記旗號的個數
func (board *Board) GetRemainingFlags() int {
	return board.remainingFlags
}

// GetElapsedTime - 取出從 startTime 之後到目前為止的時間
func (g *Game) GetElapsedTime() int {
	return int(time.Since(g.startTime).Seconds())
}
