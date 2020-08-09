package main

import (
	"fmt"
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/inkyblackness/imgui-go/v2"
	imgui2 "github.com/kyeett/guigi"
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
	imgui2.UpdatePressedCharacters()
	g.manager.Update(float32(0.1), windowWidth, windowHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	imgui2.NewFrame()

	if imgui2.CollapsingHeader("experiment") {
		if imgui2.Button("inside experiment") {
			ebitenutil.DebugPrint(screen, "Button clicked")
		}
	}

	if imgui2.UiInputText("some label2", &g.text1) {
		fmt.Println("label2 changed", g.text1)
	}
	if imgui2.Button("button") {
		ebitenutil.DebugPrint(screen, "Button clicked")
	}

	if imgui2.DragFloat("my float", &g.floatVar64) {
		ebitenutil.DebugPrint(screen, "Slider moved")
	}

	imgui2.BeginListBox("my list box")
	if imgui2.Selectable("pulsating dot") {
		fmt.Println("pulsating dot selected")
	}
	if imgui2.Selectable("fire") {
		fmt.Println("fire selected")
	}
	imgui2.EndListBox()
	imgui2.EndFrame(screen)

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
