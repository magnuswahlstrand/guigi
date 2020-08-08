package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	widgets2 "github.com/kyeett/gooigi/cmd/widgets"
	"github.com/peterhellberg/gfx"
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

var widgets []widgets2.Widget

func resetX() {
	x = pLeft + 300
}

func resetY() {
	y = pTop + 100
}

func newFrame() {
	resetX()
	resetY()

	deleteTimer--
	blinkingTimer++

	nextNoNewLine = false
	nextNoPaddingY = false

	currentListBox = ""
}

func endFrame(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v", mouse.current), 0, 20)

	for _, w := range widgets {
		w.Draw(screen)
	}
	widgets = []widgets2.Widget{}
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
