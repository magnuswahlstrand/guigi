package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/gooigi/cmd/text"
	"github.com/peterhellberg/gfx"
)

type Button struct {
	Label string
	Rect  gfx.Rect

	Pressed bool
	Hovered bool
}

func (w *Button) Draw(screen *ebiten.Image) {
	switch {
	case w.Pressed:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorPressedBlue)
	case w.Hovered:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorBlue)
	}

	text.DrawAt(screen, w.Label, text.DefaultFont, w.Rect.Min.X, w.Rect.Min.Y)
}
