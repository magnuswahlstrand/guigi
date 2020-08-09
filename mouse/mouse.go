package mouse

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/peterhellberg/gfx"
)

var mse = Mouse{}

type Mouse struct {
	pressed      bool
	dragged      bool
	justPressed  bool
	justReleased bool
	start        gfx.Vec
	previous     gfx.Vec
	current      gfx.Vec
}

func Update() {
	mse.Update()
}

func (m *Mouse) Update() {
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

func Released() bool {
	return mse.released()
}

func JustReleased() bool {
	return mse.justReleased
}

func JustPressed() bool {
	return mse.justPressed
}

func (m *Mouse) released() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

func DiffToCurrent() gfx.Vec {
	return mse.diffToCurrent()
}

func (m *Mouse) diffToCurrent() gfx.Vec {
	return m.current.Sub(m.start)
}

func DiffToPrevious() gfx.Vec {
	return mse.diffToPrevious()
}

func (m *Mouse) diffToPrevious() gfx.Vec {
	return m.previous.Sub(m.start)
}

func MouseState(r gfx.Rect) (bool, bool, bool, bool) {
	startedIn := r.Contains(mse.start)
	over := r.Contains(mse.current)
	isPressed := over && mse.pressed && startedIn
	isHovered := (over && !mse.pressed) || isPressed
	return over, startedIn, isPressed, isHovered
}

// TODO: rename
type mouseStateStruct struct {
	gfx.Rect
}

func (s *mouseStateStruct) StartedIn() bool {
	return s.Contains(mse.start)
}

func (s *mouseStateStruct) Over() bool {
	return s.Contains(mse.current)
}

func (s *mouseStateStruct) Pressed() bool {
	return mse.pressed && s.Over() && s.StartedIn()
}

func (s *mouseStateStruct) Hovered() bool {
	return s.Over() && (!mse.pressed) || (mse.pressed && s.StartedIn())
}

func (s *mouseStateStruct) Up() bool {
	return mse.justReleased && s.Over() && s.StartedIn()
}

func (s *mouseStateStruct) UpOutside() bool {
	return mse.justReleased && !s.Over() && !s.StartedIn()
}

func Dragged() bool {
	return mse.dragged
}

func (s *mouseStateStruct) Dragged() bool {
	return mse.dragged
}

func MouseStateRect(r gfx.Rect) mouseStateStruct {
	return mouseStateStruct{
		Rect: r,
	}
}
