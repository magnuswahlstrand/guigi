package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/gooigi/cmd/text"
)

var _ Widget = &Label{}

type Label struct {
	Label string
	X, Y  float64
}

func (w *Label) Draw(screen *ebiten.Image) {
	text.DrawAt(screen, w.Label, text.DefaultFont, w.X, w.Y)
}
