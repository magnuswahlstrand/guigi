package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/font"
	"image/color"
)

const (
	tOffsetX = 4
	tOffsetY = 14
)

func textDrawAt(screen *ebiten.Image, s string, fnt font.Face, x, y float64) {
	text.Draw(screen, s, fnt, int(x+tOffsetX), int(y+tOffsetY), color.White)
}

func textDrawInRectCenter(screen *ebiten.Image, s string, fnt font.Face, r gfx.Rect) {
	m := font.MeasureString(fnt, s)
	text.Draw(screen, s, fnt, int(r.Min.X+(r.W()-float64(m.Ceil()))/2), int(r.Min.Y+tOffsetY), color.White)
}
