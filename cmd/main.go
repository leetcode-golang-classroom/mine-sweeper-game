package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/game"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.DefaultScreenWidth, layout.DefaultScreenHeight)
	ebiten.SetWindowTitle(fmt.Sprintf("%s Mine Sweeper Grid", layout.LevelMessage[layout.Easy]))
	gameInstance := game.NewGame(layout.DefaultRows, layout.DefaultCols, layout.DefaultMineCounts)
	gameLayout := layout.NewGameLayout(gameInstance)
	if err := ebiten.RunGame(gameLayout); err != nil {
		log.Fatal(err)
	}
}
