package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/guigi/constants"
	"github.com/kyeett/guigi/text"
	"github.com/peterhellberg/gfx"
)

var _ Widget = &Selectable{}

type Selectable struct {
	Label string
	Rect  gfx.Rect

	Selected bool
	Hovered  bool
}

func (w *Selectable) Draw(screen *ebiten.Image) {
	switch {
	case w.Hovered:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorHoveredBlue)
	case w.Selected:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorPressedBlue)
	default:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
	}

	text.DrawAt(screen, w.Label, text.DefaultFont, w.Rect.Min.X, w.Rect.Min.Y)
}
