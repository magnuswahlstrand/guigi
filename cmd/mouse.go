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

func mouseState(r gfx.Rect) (bool, bool, bool, bool) {
	startedIn := r.Contains(mouse.start)
	over := r.Contains(mouse.current)
	isPressed := over && mouse.pressed && startedIn
	isHovered := (over && !mouse.pressed) || isPressed
	return over, startedIn, isPressed, isHovered
}

// TODO: rename
type mouseStateStruct struct {
	gfx.Rect
}

func (m *mouseStateStruct) startedIn() bool {
	return m.Contains(mouse.start)
}

func (m *mouseStateStruct) over() bool {
	return m.Contains(mouse.current)
}

func (m *mouseStateStruct) pressed() bool {
	return mouse.pressed && m.over() && m.startedIn()
}

func (m *mouseStateStruct) hovered() bool {
	return m.over() && (!mouse.pressed) || (mouse.pressed && m.startedIn())
}

func (m *mouseStateStruct) up() bool {
	return mouse.justReleased && m.over() && m.startedIn()
}

func (m *mouseStateStruct) upOutside() bool {
	return mouse.justReleased && !m.over() && !m.startedIn()
}

func mouseStateRect(r gfx.Rect) mouseStateStruct {
	return mouseStateStruct{
		Rect: r,
	}
}
