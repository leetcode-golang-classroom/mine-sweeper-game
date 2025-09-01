package layout

import "image/color"

type Level int

const (
	Easy Level = iota
	Medium
	Hard
)

type LevelSetup struct {
	Rows       int
	Cols       int
	MineCounts int
}

var LevelIconMap map[Level]string = map[Level]string{
	Easy:   "🌱",
	Medium: "⏳",
	Hard:   "💣",
}

var LevelSetupMap map[Level]LevelSetup = map[Level]LevelSetup{
	Easy: LevelSetup{
		9,
		9,
		10,
	},
	Medium: LevelSetup{
		16,
		16,
		40,
	},
	Hard: LevelSetup{
		30,
		16,
		99,
	},
}

var LevelColorMap map[Level]color.RGBA = map[Level]color.RGBA{
	Easy:   color.RGBA{0, 180, 0, 255},
	Medium: color.RGBA{39, 80, 245, 240},
	Hard:   color.RGBA{240, 0, 200, 255},
}

var LevelMessage map[Level]string = map[Level]string{
	Easy:   "Easy",
	Medium: "Medium",
	Hard:   "Hard",
}
