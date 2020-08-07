package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	text2 "github.com/kyeett/games/util/text"
	"github.com/peterhellberg/gfx"
)

type button struct {
	label string
	rect  gfx.Rect

	pressed bool
	hovered bool
}

func uiButton(label string) bool {
	bb := text2.BoundingBoxFromString(label, usedFont)
	r := rect(x, y, bb.W()+8, wHeight)
	containsStart := r.Contains(mouse.start)
	pressed := mouse.pressed && r.Contains(mouse.current)
	widgets = append(widgets, &button{
		label: label,
		rect:  r,

		pressed: containsStart && pressed,
		hovered: r.Contains(mouse.current) && (!pressed || (pressed && containsStart)),
	})

	nextLine()
	return mouse.justReleased && r.Contains(mouse.current) && containsStart
}

func (w *button) Draw(screen *ebiten.Image) {
	switch {
	case w.pressed:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorPressedBlue)
	case w.hovered:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorBlue)
	}

	textDrawAt(screen, w.label, usedFont, w.rect.Min.X, w.rect.Min.Y)
}
