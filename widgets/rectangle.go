package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/guigi/constants"
	"github.com/peterhellberg/gfx"
)

var _ Widget = &Rectangle{}

type Rectangle struct {
	Rect gfx.Rect
}

func (w *Rectangle) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
}
