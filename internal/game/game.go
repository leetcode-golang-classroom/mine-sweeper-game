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
	RemainingFlags       int              // 剩餘標記數
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
	board := NewBoard(rows, cols, mineCount)
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
func NewBoard(rows, cols, mineCount int) *Board {
	board := &Board{
		rows:                 rows,
		cols:                 cols,
		minePositionShuffler: defaultPositionShuffler,
		RemainingFlags:       mineCount,
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

// ToggleFlag - 標記地雷
func (board *Board) ToggleFlag(row, col int) {
	// 超出邊界
	if row < 0 || row >= board.rows ||
		col < 0 || col >= board.cols {
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
	if board.RemainingFlags-count < 0 {
		return
	}
	board.RemainingFlags -= count
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
		if curRow < 0 || curRow >= board.rows ||
			curCol < 0 || curCol >= board.cols {
			continue
		}

		cell := board.cells[curRow][curCol]
		// 已經被揭開
		if cell.Revealed {
			continue
		}

		// 標注該格已經被揭開
		board.cells[curRow][curCol].Revealed = true
		if cell.Flagged {
			board.RemainingFlags++
			board.cells[curRow][curCol].Flagged = false
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
