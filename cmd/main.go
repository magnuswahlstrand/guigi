package main

import (
	"fmt"
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

	text1, text2 string
	c1           [4]float32
	c2           [4]float32
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (g *Game) Update(_ *ebiten.Image) error {
	mouse.update()
	updatePressedCharacters()
	g.manager.Update(float32(0.1), windowWidth, windowHeight)
	return nil
}

var keysPressed string

func (g *Game) Draw(screen *ebiten.Image) {
	newFrame()

	if UiInputText("some label1", &g.text1) {
		fmt.Println("label1 changed", g.text1)
	}
	if UiInputText("some label2", &g.text2) {
		fmt.Println("label2 changed", g.text2)
	}
	if UiButton("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if UiButton("button2") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}
	if UiDragFloat("my float", &g.floatVar64) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}
	endFrame(screen)

	g.manager.BeginFrame()

	f1 := func() {
		if imgui.InputText("some label1", &g.text1) {
			fmt.Println("label1 changed", g.text1)
		}
		//imgui.ColorEdit4("color1", &g.c1)
	}

	f2 := func() {
		if imgui.InputText("some label2", &g.text2) {
			fmt.Println("label2 changed")
		}
		//imgui.ColorEdit4("color2", &g.c2)
	}

	f1()
	f2()

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

func GetRect() gfx.Rect {
	return gfx.R(0, 0, wWidth, wHeight).Moved(gfx.V(x, y))
}

func GetRectWide(w float64) gfx.Rect {
	return gfx.R(0, 0, w, wHeight).Moved(gfx.V(x, y))
}

func GetXY() (float64, float64) {
	return x, y
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

		text1: "magnus",
		text2: "rui",

		c1: [4]float32{1, 1, 1, 1},
		c2: [4]float32{1, 1, 1, 1},
	}

	ebiten.RunGame(g)
}
