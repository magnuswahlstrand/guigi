package main

import (
	"fmt"
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/inkyblackness/imgui-go/v2"
)

var (
	x, y  float64
	mouse Mouse
)

type Game struct {
	manager      *renderer.Manager
	floatVar     float32
	floatVar64   float64
	exampleIndex int32

	text1 string
	c1    [4]float32
	c2    [4]float32
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

	if UiCollapsingHeader("experiment") {
		if UiButton("inside experiment") {
			ebitenutil.DebugPrint(screen, "Button clicked")
		}
	}

	if UiInputText("some label2", &g.text1) {
		fmt.Println("label2 changed", g.text1)
	}
	if UiButton("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}

	if UiDragFloat("my float", &g.floatVar64) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}

	UiBeginListBox("my list box")
	if UiSelectable("pulsating dot") {
		fmt.Println("pulsating dot selected")
	}
	if UiSelectable("fire") {
		fmt.Println("fire selected")
	}
	UiEndListBox()

	endFrame(screen)

	g.manager.BeginFrame()
	if imgui.CollapsingHeader("experiment") {
		if imgui.Button("inside experiment") {
			ebitenutil.DebugPrint(screen, "Button clicked")
		}
	}

	if imgui.InputText("some label2", &g.text1) {
		fmt.Println("label2 changed")
	}

	if imgui.Button("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}

	if imgui.DragFloat("my float", &g.floatVar) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}

	if imgui.ListBox("", &g.exampleIndex, []string{"pulsating dot", "fire"}) {
		switch g.exampleIndex {
		case 0:
			fmt.Println("Pulsating")
		case 1:
			fmt.Println("Fire")
		}
	}
	g.manager.EndFrame(screen)
}

const (
	windowWidth  = 800
	windowHeight = 600
)

const (
	pTop  = 20 + 10
	pLeft = 10
	// TODO: Used during prototyping, should be removed
	devOffsetX = 300
	devOffsetY = 100
	frameWidth = 250
	wHeight    = 20
	wWidth     = 100

	wPaddingY = 5
	wPaddingX = 1
)

func main() {
	mgr := renderer.New(nil)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	g := &Game{
		manager:    mgr,
		floatVar:   0.0,
		floatVar64: 0.0,

		text1: "magnus",

		c1: [4]float32{1, 1, 1, 1},
		c2: [4]float32{1, 1, 1, 1},
	}

	ebiten.RunGame(g)
}
