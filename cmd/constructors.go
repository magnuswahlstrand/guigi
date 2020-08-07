package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	text2 "github.com/kyeett/games/util/text"
	widgets2 "github.com/kyeett/gooigi/cmd/widgets"
)

func UiInputText(label string, variable *string) bool {
	r := GetRect()

	// Check if widget clicked this turn
	var updated bool
	if mouse.pressed && r.Contains(mouse.current) {
		setActiveLabel(label)
	}

	// If widget is current active, handle key presses
	active := label == activeLabel
	if active {
		updated = tryUpdateInput(variable)
	}

	w := &widgets2.InputText{
		Label:    label,
		Rect:     r,
		Variable: variable,
		Active:   active,
	}
	widgets = append(widgets, w)

	nextLine()

	return updated
}

func UiDragFloat(label string, v *float64) bool {
	stepSize := 0.1
	r := GetRect()
	containsStart := r.Contains(mouse.start)

	active := mouse.pressed && containsStart
	widgets = append(widgets, &widgets2.DragFloat{
		Label: label,
		Rect:  r,

		Variable: v,
		Format:   "%0.2f",

		Active:  active,
		Hovered: r.Contains(mouse.current) && (!mouse.pressed || (mouse.pressed && containsStart)),
	})

	if active && mouse.dragged {
		diff := mouse.diffToCurrent().X - mouse.diffToPrevious().X
		*v += stepSize * diff
	}

	// Add label
	x += r.W() + wPaddingX
	widgets = append(widgets, &widgets2.Label{
		Label: label,
		X:     x,
		Y:     y,
	})

	// Move down one line
	resetX()
	y += wHeight + wPaddingY

	return active && mouse.dragged
}

func UiButton(label string) bool {
	bb := text2.BoundingBoxFromString(label, usedFont)
	r := GetRectWide(bb.W() + 8)
	containsStart := r.Contains(mouse.start)
	pressed := mouse.pressed && r.Contains(mouse.current)
	widgets = append(widgets, &widgets2.Button{
		Label: label,
		Rect:  r,

		Pressed: containsStart && pressed,
		Hovered: r.Contains(mouse.current) && (!pressed || (pressed && containsStart)),
	})

	nextLine()
	return mouse.justReleased && r.Contains(mouse.current) && containsStart
}

func tryUpdateInput(variable *string) bool {
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		// String already empty
		if len(*variable) == 0 {
			return false
		}

		// Last delete happened to recently
		if deleteTimer > 0 {
			return false
		}

		newTimer := deleteTimerMin
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			newTimer = deleteTimerMax
		}

		*variable = (*variable)[:len(*variable)-1]
		deleteTimer = newTimer

		return true
	}

	if len(keysPressed) > 0 {
		*variable += keysPressed
		return true
	}
	return false
}
