package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	text2 "github.com/kyeett/games/util/text"
	widgets2 "github.com/kyeett/gooigi/cmd/widgets"
	"github.com/peterhellberg/gfx"
)

func UiInputText(label string, variable *string) bool {
	r := AllocateRect()
	_, _, isPressed, _ := mouseState(r)

	// Check if widget clicked this turn
	if isPressed {
		setFocused(label)
	}

	// If widget is current focused, handle key presses
	focused := isFocused(label)
	var updated bool
	if focused {
		updated = tryUpdateInput(variable)
	}

	w := &widgets2.InputText{
		Label:    label,
		Rect:     r,
		Variable: variable,
		Focused:  focused,

		ShowBlinker: focused && shouldShowBlinker(),
	}
	sameLine()
	addWidget(w)

	// Add label
	x, y := AllocateXY()
	l := &widgets2.Label{
		Label: label,
		X:     x,
		Y:     y,
	}
	addWidget(l)
	return updated
}

func shouldShowBlinker() bool {
	// Last delete happened too recently
	return blinkingTimer/30%2 == 0
}

func mouseState(r gfx.Rect) (bool, bool, bool, bool) {
	startedIn := r.Contains(mouse.start)
	over := r.Contains(mouse.current)
	isPressed := over && mouse.pressed && startedIn
	isHovered := (over && !mouse.pressed) || isPressed
	return over, startedIn, isPressed, isHovered
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

func UiDragFloat(label string, v *float64) bool {
	stepSize := 0.1
	r := AllocateRect()

	_, startedIn, _, isHovered := mouseState(r)
	dragged := mouse.pressed && startedIn

	// Update if needed
	if dragged && mouse.dragged {
		diff := mouse.diffToCurrent().X - mouse.diffToPrevious().X
		*v += stepSize * diff
	}

	w := &widgets2.DragFloat{
		Label: label,
		Rect:  r,

		Variable: v,
		Format:   "%0.2f",

		Active:  dragged,
		Hovered: isHovered,
	}
	sameLine()
	addWidget(w)

	// Add label
	x, y := AllocateXY()
	l := &widgets2.Label{
		Label: label,
		X:     x + wPaddingX,
		Y:     y,
	}
	addWidget(l)

	return dragged && mouse.dragged
}

func sameLine() {
	keepSameLine = true
}

func UiButton(label string) bool {
	bb := text2.BoundingBoxFromString(label, usedFont)
	r := AllocateWideRect(bb.W() + 8)
	over, startedIn, isPressed, isHovered := mouseState(r)

	w := &widgets2.Button{
		Label: label,
		Rect:  r,

		Pressed: isPressed,
		Hovered: isHovered,
	}
	addWidget(w)

	return mouse.justReleased && over && startedIn
}

func UiCollapsingHeader(label string) bool {
	r := AllocateRect()
	over, startedIn, _, isHovered := mouseState(r)

	// Update state if just clicked
	expanded := isActive("CollapsingHeader", label)

	mouseUp := mouse.justReleased && over && startedIn
	if mouseUp {
		if expanded {
			expanded = false
		} else {
			expanded = true
		}

		setActive("CollapsingHeader", label, expanded)
	}

	w := &widgets2.CollapsingHeader{
		Label:   label,
		Rect:    r,
		Hovered: isHovered,
	}
	addWidget(w)

	return expanded
}
