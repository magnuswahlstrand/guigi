package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/gooigi/cmd/text"
	"github.com/peterhellberg/gfx"
)

type InputText struct {
	Label string
	Rect  gfx.Rect

	Variable *string

	Active bool
}

func (w *InputText) Draw(screen *ebiten.Image) {
	if w.Active {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorHoveredBlue)
	} else {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorBlue)
	}

	text.DrawAt(screen, *w.Variable, text.DefaultFont, w.Rect.Min.X, w.Rect.Min.Y)
}
