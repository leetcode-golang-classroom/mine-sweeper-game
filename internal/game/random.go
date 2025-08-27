package game

import (
	"math/rand"
	"time"
)

var defaultPositionShuffler positionShuffler = func(coords []coord) {
	if len(coords) <= 1 {
		return
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(coords), func(i, j int) {
		coords[i], coords[j] = coords[j], coords[i]
	})
}
