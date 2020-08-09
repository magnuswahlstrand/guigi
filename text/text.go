package text

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

const (
	tOffsetX = 4
	tOffsetY = 14
)

var (
	DefaultFont font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	DefaultFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func DrawAt(screen *ebiten.Image, s string, fnt font.Face, x, y float64) {
	text.Draw(screen, s, fnt, int(x+tOffsetX), int(y+tOffsetY), color.White)
}

func DrawInRectCenter(screen *ebiten.Image, s string, fnt font.Face, r gfx.Rect) {
	m := font.MeasureString(fnt, s)
	text.Draw(screen, s, fnt, int(r.Min.X+(r.W()-float64(m.Ceil()))/2), int(r.Min.Y+tOffsetY), color.White)
}
