package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/util"
	"github.com/kyeett/guigi/constants"
	"github.com/peterhellberg/gfx"
)

var (
	checkmarkImage *ebiten.Image
)

func init() {
	checkmarkImage = util.LoadAssetImageOrFatal(Asset, "checkbox.png")
}

type Checkbox struct {
	Label string
	Rect  gfx.Rect

	Checked bool
	Hovered bool
}

func (w *Checkbox) Draw(screen *ebiten.Image) {
	if w.Hovered {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorHoveredBlue)
	} else {
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
	}

	if w.Checked {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(w.Rect.Min.X+1, w.Rect.Min.Y+3)
		screen.DrawImage(checkmarkImage, opt)
	}
}
