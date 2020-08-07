package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
)

type dragFloat struct {
	label string
	rect  gfx.Rect

	variable *float64
	format   string

	active  bool
	hovered bool
}

func (w *dragFloat) Draw(screen *ebiten.Image) {
	switch {
	case w.active:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorPressedBlue)
	case w.hovered:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorHoveredBlue)
	default:
		ebitenutil.DrawRect(screen, w.rect.Min.X, w.rect.Min.Y, w.rect.W(), w.rect.H(), colorBlue)
	}

	s := fmt.Sprintf(w.format, *w.variable)
	textDrawInRectCenter(screen, s, usedFont, w.rect)
}

func uiDragFloat(label string, v *float64) bool {
	stepSize := 0.1
	r := rect(x, y, wWidth, wHeight)
	containsStart := r.Contains(mouse.start)

	active := mouse.pressed && containsStart
	widgets = append(widgets, &dragFloat{
		label: label,
		rect:  r,

		variable: v,
		format:   "%0.2f",

		active:  active,
		hovered: r.Contains(mouse.current) && (!mouse.pressed || (mouse.pressed && containsStart)),
	})

	if active && mouse.dragged {
		diff := mouse.diffToCurrent().X - mouse.diffToPrevious().X
		*v += stepSize * diff
	}

	// Add label
	x += r.W() + wPaddingX
	widgets = append(widgets, &Label{
		label: label,
		X:     x,
		Y:     y,
	})

	// Move down one line
	resetX()
	y += wHeight + wPaddingY

	return active && mouse.dragged
}
