package main

import (
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/inkyblackness/imgui-go/v2"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

var (
	usedFont font.Face

	x, y             float64
	mouse            Mouse
	colorBlue        color.Color
	colorHoveredBlue color.Color
	colorPressedBlue color.Color
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	usedFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type Game struct {
	manager    *renderer.Manager
	floatVar   float32
	floatVar64 float64
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (g *Game) Update(_ *ebiten.Image) error {
	mouse.update()
	g.manager.Update(float32(0.1), windowWidth, windowHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	newFrame()
	if uiButton("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if uiButton("button2") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if uiDragFloat("my float", &g.floatVar64) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}
	endFrame(screen)

	g.manager.BeginFrame()
	if imgui.Button("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if imgui.Button("button2") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if imgui.DragFloat("my float", &g.floatVar) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}
	g.manager.EndFrame(screen)
}

const (
	windowWidth  = 800
	windowHeight = 600
)

const (
	pTop    = 10
	pLeft   = 10
	wHeight = 20
	wWidth  = 100

	wPaddingY = 5
	wPaddingX = 1
)

func rect(x, y, w, h float64) gfx.Rect {
	return gfx.R(0, 0, w, h).Moved(gfx.V(x, y))
}

type Widget interface {
	Draw(screen *ebiten.Image)
}

func main() {
	colorBlue = color.RGBA{26, 66, 114, 255}
	colorPressedBlue = color.RGBA{38, 130, 255, 255}
	colorHoveredBlue = color.RGBA{0, 140, 251, 255}

	mgr := renderer.New(nil)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	g := &Game{
		manager:    mgr,
		floatVar:   0.0,
		floatVar64: 0.0,
	}

	ebiten.RunGame(g)
}
