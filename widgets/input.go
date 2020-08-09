package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	text2 "github.com/kyeett/games/util/text"
	"github.com/kyeett/guigi/constants"
	"github.com/kyeett/guigi/text"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var _ Widget = &InputText{}

type InputText struct {
	Label string
	Rect  gfx.Rect

	Variable *string

	Focused     bool
	ShowBlinker bool
}

func (w *InputText) Draw(screen *ebiten.Image) {
	if w.Focused {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorHoveredBlue)
	} else {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
	}

	text.DrawAt(screen, *w.Variable, text.DefaultFont, w.Rect.Min.X, w.Rect.Min.Y)

	if w.ShowBlinker {
		bb := text2.BoundingBoxFromString(*w.Variable, text.DefaultFont)
		// TODO: remove hardcoded values
		ebitenutil.DrawRect(screen, w.Rect.Min.X+bb.W()+6, w.Rect.Min.Y+4, 1, bb.H(), colornames.White)
	}

}
