package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/guigi/constants"
	"github.com/kyeett/guigi/text"
	"github.com/peterhellberg/gfx"
)

var _ Widget = &Button{}

type Button struct {
	Label string
	Rect  gfx.Rect

	Pressed bool
	Hovered bool
}

func (w *Button) Draw(screen *ebiten.Image) {
	switch {
	case w.Pressed:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorPressedBlue)
	case w.Hovered:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
	}

	text.DrawAt(screen, w.Label, text.DefaultFont, w.Rect.Min.X, w.Rect.Min.Y)
}
