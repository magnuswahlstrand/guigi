package widgets

import (
	"github.com/hajimehoshi/ebiten"
)

type Widget interface {
	Draw(screen *ebiten.Image)
}
