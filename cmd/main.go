package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/game"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.ScreenWidth, layout.ScreenHeight)
	ebiten.SetWindowTitle("Mine Sweeper Grid")
	gameInstance := game.NewGame(layout.Rows, layout.Cols, layout.MineCounts)
	gameLayout := layout.NewGameLayout(gameInstance)
	if err := ebiten.RunGame(gameLayout); err != nil {
		log.Fatal(err)
	}
}
