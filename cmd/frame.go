package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func nextLine() {
	y += wHeight + wPaddingY
}

var widgets []Widget

func resetX() {
	x = pLeft + 300
}

func resetY() {
	y = pTop + 100
}

func newFrame() {
	resetX()
	resetY()
}

func endFrame(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v", mouse.current), 0, 20)

	for _, w := range widgets {
		w.Draw(screen)
	}
	widgets = []Widget{}
}
