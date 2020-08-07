package main

import "github.com/hajimehoshi/ebiten"

type Label struct {
	label string
	X, Y  float64
}

func (w *Label) Draw(screen *ebiten.Image) {
	textDrawAt(screen, w.label, usedFont, w.X, w.Y)
}
