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

	keepSameLine bool

	activeList = map[string]bool{}
)

func setFocused(label string) {
	if focusedLabel != label {
		// New widget
		focusedLabel = label

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
