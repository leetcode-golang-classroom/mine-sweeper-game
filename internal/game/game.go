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
	rows  int       // 總共格數
	cols  int       // 總共列數
	cells [][]*Cell // 整格棋盤狀態
}

// Game - 遊戲物件
type Game struct {
	board       *Board // 棋盤物件
	isGameOver  bool   // 是否遊戲結束
	isPlayerWin bool   // 玩家是否獲勝
	remaining   int    // 剩餘需要翻開的格子數
}

func NewGame(rows, cols, mineNumbers int) *Game {
	board := createBoard(rows, cols)
	return &Game{
		board:       board,
		remaining:   rows*cols - mineNumbers,
		isGameOver:  false,
		isPlayerWin: false,
	}
}

func createBoard(rows, cols int) *Board {
	board := &Board{
		rows: rows,
		cols: cols,
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
func (g *Game) Init(board *Board) {
	// 無效的設定
	if board == nil || len(board.cells) != board.rows || len(board.cells[0]) != board.cols {
		return
	}
	// 設定資料
	for row := range board.cells {
		for col := range board.cells[row] {
			sourceCell := board.cells[row][col]
			g.board.cells[row][col].AdjacenetMines = sourceCell.AdjacenetMines
			g.board.cells[row][col].IsMine = sourceCell.IsMine
			g.board.cells[row][col].Revealed = sourceCell.Revealed
			g.board.cells[row][col].Flagged = sourceCell.Flagged
		}
	}
}
