package main

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	focusedLabel = ""

	blinkingTimer = 0

	deleteTimer    = 0
	deleteTimerMax = 10
	deleteTimerMin = 4

	nextNoNewLine  bool
	nextNoPaddingY bool

	activeList   = map[string]bool{}
	selectedList = map[string]int{}

	currentListBox      = ""
	currentListBoxIndex = -1
)

func getCurrentIndex() int {
	currentListBoxIndex++
	return currentListBoxIndex
}

func setFocused(label string) {
	if focusedLabel != label {
		// New widget
		focusedLabel = label

		// Reset timers
		blinkingTimer = 0
		deleteTimer = 0
	}
}

func setUnfocused(label string) {
	if focusedLabel == label {
		focusedLabel = ""

		// Reset timers
		blinkingTimer = 0
		deleteTimer = 0
	}
}

func isFocused(label string) bool {
	return label == focusedLabel
}

func updatePressedCharacters() {
	keysPressed = ""
	for _, r := range ebiten.InputChars() {
		keysPressed += string(r)
	}
}

func setActive(typ string, label string, state bool) {
	activeList[typ+label] = state
}

func isActive(typ string, label string) bool {
	return activeList[typ+label]
}

func setSelectedIndex(typ string, label string, index int) {
	selectedList[typ+label] = index
}

func isSelectedIndex(typ string, label string, index int) bool {
	currentIndex, found := selectedList[typ+label]
	return found && (currentIndex == index)
}

func sameLine() {
	nextNoNewLine = true
}

func noPaddingY() {
	nextNoPaddingY = true
}
