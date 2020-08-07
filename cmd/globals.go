package main

import "github.com/hajimehoshi/ebiten"

var (
	activeLabel = ""

	blinkingTimer    = 0
	blinkingTimerMax = 5

	deleteTimer    = 0
	deleteTimerMax = 10
	deleteTimerMin = 4
)

func setActiveLabel(label string) {
	if activeLabel != label {
		// New widget
		activeLabel = label

		// Reset timers
		blinkingTimer = 0
		deleteTimer = 0
	}
}

func updatePressedCharacters() {
	keysPressed = ""
	for _, r := range ebiten.InputChars() {
		keysPressed += string(r)
	}
}
