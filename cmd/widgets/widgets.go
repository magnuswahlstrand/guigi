package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Widget interface {
	Draw(screen *ebiten.Image)
}

var (
	colorBlue        = color.RGBA{26, 66, 114, 255}
	colorPressedBlue = color.RGBA{38, 130, 255, 255}
	colorHoveredBlue = color.RGBA{0, 140, 251, 255}
)
