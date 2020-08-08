package widgets

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/gooigi/cmd/text"
	"github.com/peterhellberg/gfx"
)

var _ Widget = &DragFloat{}

type DragFloat struct {
	Label string
	Rect  gfx.Rect

	Variable *float64
	Format   string

	Active  bool
	Hovered bool
}

func (w *DragFloat) Draw(screen *ebiten.Image) {
	switch {
	case w.Active:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorPressedBlue)
	case w.Hovered:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), colorBlue)
	}

	s := fmt.Sprintf(w.Format, *w.Variable)
	text.DrawInRectCenter(screen, s, text.DefaultFont, w.Rect)
}
