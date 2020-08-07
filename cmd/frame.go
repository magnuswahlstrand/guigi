package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	widgets2 "github.com/kyeett/gooigi/cmd/widgets"
)

func nextLine() {
	y += wHeight + wPaddingY
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
	blinkingTimer--

	fmt.Println(deleteTimer, blinkingTimer)
}

func endFrame(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v", mouse.current), 0, 20)

	for _, w := range widgets {
		w.Draw(screen)
	}
	widgets = []widgets2.Widget{}
}
