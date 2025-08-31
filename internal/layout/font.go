package layout

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leetcode-golang-classroom/mine-sweeper/internal/fonts"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	// emojiFace       font.Face
	emojiFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	s, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.NotoEmojiRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	emojiFaceSource = s
}

const (
	IsMine = iota + 100
	IsFlag
	IsButtonIcon
)

func getTileColor(value int) color.Color {
	switch value {
	case 0:
		return color.RGBA{0x77, 0x6e, 0x65, 0xff}
	case IsFlag:
		return color.RGBA{0xf9, 0xf6, 0xf2, 0xff}
	case IsMine, IsButtonIcon:
		return color.Black
	default:
		return color.RGBA{0xf9, 0xf6, 0xf2, 0xff}
	}
}
