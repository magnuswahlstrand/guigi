package guigi

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/guigi/mouse"
	widgets2 "github.com/kyeett/guigi/widgets"
	"github.com/peterhellberg/gfx"
	"log"
)

const (
	pTop  = 20 + 10
	pLeft = 10
	// TODO: Used during prototyping, should be removed
	devOffsetX = 300
	devOffsetY = 100
	frameWidth = 250
	wHeight    = 20
	wWidth     = 100

	wPaddingY = 5
	wPaddingX = 1
)

func nextLine() {
	resetX()

	y += wHeight
	if nextNoPaddingY {
		nextNoPaddingY = false
	} else {
		y += wPaddingY
	}
}

func resetX() {
	x = pLeft + devOffsetX
}

var widgets []widgets2.Widget

type frameDimensions struct {
	x, y float64
	w, h float64
}

func (d *frameDimensions) colX(col int) float64 {
	switch col {
	case 0:
		return 0
	case 1:
		return cw
	case 2:
		return d.w - cw
	}
	log.Fatal("invalid col")
	return -1
}

func (d *frameDimensions) rowY(row int) float64 {
	switch row {
	case 0:
		return 0
	case 1:
		return ch
	case 2:
		return d.h - ch
	}
	log.Fatal("invalid row")
	return -1
}

func (d *frameDimensions) translate(row, col int) *ebiten.DrawImageOptions {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(f.x, f.y)

	opt.GeoM.Translate(d.colX(col), d.rowY(row))
	return opt
}

func (d *frameDimensions) translateExisting(opt *ebiten.DrawImageOptions, row, col int) *ebiten.DrawImageOptions {
	opt.GeoM.Translate(f.x, f.y)
	opt.GeoM.Translate(d.colX(col), d.rowY(row))
	return opt
}

func (d *frameDimensions) midWidth() float64 {
	return d.w - cw*2
}

func (d *frameDimensions) midHeight() float64 {
	return d.h - ch*2
}

var f frameDimensions

func NewFrame() {
	x = devOffsetX
	y = devOffsetY
	f.x = x
	f.y = y

	x += pLeft
	y += pTop

	deleteTimer--
	blinkingTimer++

	nextNoNewLine = false
	nextNoPaddingY = false

	currentListBox = ""

	widgets = []widgets2.Widget{}

	mouse.Update()
}

func EndFrame(screen *ebiten.Image) {
	// Draw base frame
	f.w = frameWidth
	f.h = y - f.y

	// Mid
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(f.w-2*cw, f.h-2*ch)
	screen.DrawImage(mid, f.translateExisting(opt, 1, 1))

	// Corners
	screen.DrawImage(tl, f.translate(0, 0))
	screen.DrawImage(bl, f.translate(2, 0))
	screen.DrawImage(br, f.translate(2, 2))
	screen.DrawImage(tr, f.translate(0, 2))

	// Top
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(f.midWidth(), 1)
	screen.DrawImage(t, f.translateExisting(opt, 0, 1))

	// Left
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(1, f.midHeight())
	screen.DrawImage(l, f.translateExisting(opt, 1, 0))

	// Bottom
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(f.midWidth(), 1)
	screen.DrawImage(b, f.translateExisting(opt, 2, 1))
	//
	// Right
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(1, f.midHeight())
	screen.DrawImage(r, f.translateExisting(opt, 1, 2))

	// Draw widgets
	for _, w := range widgets {
		w.Draw(screen)
	}
}

func addWidget(w widgets2.Widget) {
	widgets = append(widgets, w)

	if !nextNoNewLine {
		nextLine()
	}

	// Reset
	nextNoNewLine = false
}

func allocateRect() gfx.Rect {
	return allocateRectW(wWidth)
}

func allocateRectW(w float64) gfx.Rect {
	r := gfx.R(0, 0, w, wHeight).Moved(gfx.V(x, y))
	x += w + wPaddingX
	return r
}

func allocateRectH(h float64) gfx.Rect {
	r := gfx.R(0, 0, wWidth, h).Moved(gfx.V(x, y))
	x += wWidth + wPaddingX
	return r
}

func allocateXY() (float64, float64) {
	x += wPaddingX
	return x, y
}
