package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/peterhellberg/gfx"
)

type Mouse struct {
	pressed      bool
	dragged      bool
	justPressed  bool
	justReleased bool
	start        gfx.Vec
	previous     gfx.Vec
	current      gfx.Vec
}

func (m *Mouse) update() {
	// Update previous
	m.previous = m.current

	// Update current
	cx, cy := ebiten.CursorPosition()
	m.current = gfx.IV(cx, cy)

	m.dragged = !m.current.Eq(m.previous)

	// Update states
	m.pressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	m.justPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	if m.justPressed {
		m.start = m.current
		m.previous = m.current
	}
	m.justReleased = inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

func (m *Mouse) released() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

func (m *Mouse) diffToCurrent() gfx.Vec {
	return m.current.Sub(m.start)
}

func (m *Mouse) diffToPrevious() gfx.Vec {
	return m.previous.Sub(m.start)
}
