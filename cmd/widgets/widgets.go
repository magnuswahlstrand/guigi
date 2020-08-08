package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Widget interface {
	Draw(screen *ebiten.Image)
}

var (
	colorBlue        = color.RGBA{22, 41, 65, 255}
	colorPressedBlue = color.RGBA{21, 70, 120, 255}
	colorHoveredBlue = color.RGBA{12, 114, 203, 255}
)
