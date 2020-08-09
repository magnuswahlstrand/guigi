package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/util"
	"github.com/kyeett/gooigi/cmd/constants"
	"github.com/kyeett/gooigi/cmd/text"
	"github.com/peterhellberg/gfx"
	"math"
)

var (
	arrowImage  *ebiten.Image
	arrowWidth  float64
	arrowHeight float64
)

func init() {
	arrowImage = util.LoadAssetImageOrFatal(Asset, "arrow.png")
	arrowBounds := gfx.BoundsToRect(arrowImage.Bounds())
	arrowWidth = arrowBounds.W()
	arrowHeight = arrowBounds.H()
}

const (
	arrowPaddingLeft  = 8
	arrowPaddingRight = 4
)

var _ Widget = &CollapsingHeader{}

type CollapsingHeader struct {
	Label string
	Rect  gfx.Rect

	Hovered   bool
	Collapsed bool
}

func (w *CollapsingHeader) Draw(screen *ebiten.Image) {
	switch {
	case w.Hovered:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.Rect.Min.X, w.Rect.Min.Y, w.Rect.W(), w.Rect.H(), constants.ColorBlue)
	}

	opt := &ebiten.DrawImageOptions{}
	if !w.Collapsed {
		opt.GeoM.Translate(-arrowWidth/2, -arrowHeight/2)
		opt.GeoM.Rotate(math.Pi / 2)
		opt.GeoM.Translate(arrowWidth/2, arrowHeight/2)
	}
	opt.GeoM.Translate(w.Rect.Min.X+arrowPaddingLeft, w.Rect.Min.Y+(w.Rect.H()-arrowHeight)/2)
	screen.DrawImage(arrowImage, opt)

	text.DrawAt(screen, w.Label, text.DefaultFont, w.Rect.Min.X+arrowPaddingLeft+arrowWidth+arrowPaddingRight, w.Rect.Min.Y)
}
