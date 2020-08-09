package imgui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/util"
	"github.com/kyeett/guigi/constants"
)

var (
	tl *ebiten.Image
	tr *ebiten.Image
	br *ebiten.Image
	bl *ebiten.Image

	mid *ebiten.Image

	t *ebiten.Image
	b *ebiten.Image
	l *ebiten.Image
	r *ebiten.Image

	arrowWidth  float64
	arrowHeight float64
)

const (
	cw, ch = 18, 18
)

func init() {

	maskCorner := util.LoadAssetImageOrFatal(Asset, "mask_topleft.png")
	borderCorner := util.LoadAssetImageOrFatal(Asset, "border_topleft.png")
	borderTop := util.LoadAssetImageOrFatal(Asset, "border_top.png")
	borderLeft := util.LoadAssetImageOrFatal(Asset, "border_left.png")

	// Corners
	tl, _ = ebiten.NewImage(cw, ch, ebiten.FilterDefault)
	tr, _ = ebiten.NewImage(cw, ch, ebiten.FilterDefault)
	br, _ = ebiten.NewImage(cw, ch, ebiten.FilterDefault)
	bl, _ = ebiten.NewImage(cw, ch, ebiten.FilterDefault)

	// Generate top left
	tl.Fill(constants.ColorBackground)

	opt := &ebiten.DrawImageOptions{}
	opt.CompositeMode = ebiten.CompositeModeDestinationAtop
	tl.DrawImage(maskCorner, opt)
	opt = &ebiten.DrawImageOptions{}
	tl.DrawImage(borderCorner, opt)

	opt = &ebiten.DrawImageOptions{}
	flipLR(opt)
	tr.DrawImage(tl, opt)

	opt = &ebiten.DrawImageOptions{}
	flipLR(opt)
	flipTD(opt)
	br.DrawImage(tl, opt)

	opt = &ebiten.DrawImageOptions{}
	flipTD(opt)
	bl.DrawImage(tl, opt)

	// Mid
	mid, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	mid.Fill(constants.ColorBackground)

	// Borders
	t, _ = ebiten.NewImage(1, ch, ebiten.FilterDefault)
	t.Fill(constants.ColorBackground)
	t.DrawImage(borderTop, &ebiten.DrawImageOptions{})

	l, _ = ebiten.NewImage(cw, 1, ebiten.FilterDefault)
	l.Fill(constants.ColorBackground)
	l.DrawImage(borderLeft, &ebiten.DrawImageOptions{})

	b, _ = ebiten.NewImage(1, ch, ebiten.FilterDefault)
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(1, -1)
	opt.GeoM.Translate(0, ch)
	b.DrawImage(t, opt)

	r, _ = ebiten.NewImage(cw, ch, ebiten.FilterDefault)
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(-1, 1)
	opt.GeoM.Translate(cw, 0)
	r.DrawImage(l, opt)

	//
	//opt := &ebiten.DrawImageOptions{}
	//opt.CompositeMode = ebiten.CompositeModeDestinationAtop
	//topLeft.DrawImage(maskCorner, opt)
	//
	//flipLR(opt)
	//
	//opt.GeoM.Scale(1, -1)
}

func flipTD(opt *ebiten.DrawImageOptions) {
	opt.GeoM.Translate(-cw/2, -ch/2)
	opt.GeoM.Scale(1, -1) // Left-right
	opt.GeoM.Translate(cw/2, ch/2)
}

func flipLR(opt *ebiten.DrawImageOptions) {
	opt.GeoM.Translate(-cw/2, -ch/2)
	opt.GeoM.Scale(-1, 1) // Left-right
	opt.GeoM.Translate(cw/2, ch/2)
}
