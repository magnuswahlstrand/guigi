package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	text2 "github.com/kyeett/games/util/text"
	"github.com/kyeett/gooigi/cmd/text"
	widgets2 "github.com/kyeett/gooigi/cmd/widgets"
)

func UiInputText(label string, variable *string) bool {
	r := allocateRect()
	//_, _, isPressed, _ := mouseState(r)
	s := mouseStateRect(r)

	// Check if widget clicked this turn

	switch {
	case s.pressed():
		setFocused(label)
	case mouse.justPressed: //&& s.pressed(): This is already implicit by previous case
		setUnfocused(label)
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
	x, y := allocateXY()
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
	r := allocateRect()

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
	x, y := allocateXY()
	l := &widgets2.Label{
		Label: label,
		X:     x + wPaddingX,
		Y:     y,
	}
	addWidget(l)

	return dragged && mouse.dragged
}

func UiButton(label string) bool {
	bb := text2.BoundingBoxFromString(label, text.DefaultFont)
	r := allocateRectW(bb.W() + 8)
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
	r := allocateRect()
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
		Label:     label,
		Rect:      r,
		Hovered:   isHovered,
		Collapsed: !expanded,
	}
	addWidget(w)

	return expanded
}

func UiBeginListBox(label string) {
	currentListBox = label
	currentListBoxIndex = -1
}

func UiEndListBox() {
	// Add empty bottom rectangle
	r := allocateRectH(6)
	w := &widgets2.Rectangle{
		Rect: r,
	}
	addWidget(w)

	currentListBox = ""
	currentListBoxIndex = -1
}

func UiSelectable(itemLabel string) bool {
	index := getCurrentIndex()

	//var selected bool

	r := allocateRect()

	// Get current state
	s := mouseStateRect(r)

	// Update selected
	if s.up() {
		setSelectedIndex("ListBox", currentListBox, index)
	}

	selected := isSelectedIndex("ListBox", currentListBox, index)

	// Add widget
	w := &widgets2.Selectable{
		Label: itemLabel,
		Rect:  r,

		Selected: selected,
		Hovered:  s.hovered(),
	}
	noPaddingY()
	addWidget(w)

	return s.up()
}
