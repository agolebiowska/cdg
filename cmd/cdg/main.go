package main

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/agolebiowska/cdg/pkg/scene"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func initScreen() {
	Global.PrimaryMonitor = pixelgl.PrimaryMonitor()
	cfg := pixelgl.WindowConfig{
		Title:  Global.Title,
		Bounds: pixel.R(0, 0, Global.WindowWidth, Global.WindowHeight),
		VSync:  Global.Vsync,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	Global.Win = window
}

func handleInput() {
	if Global.Win.JustPressed(pixelgl.KeyEscape) {
		Global.Win.SetClosed(true)
	}
	// control the actor with keys
	if Global.Win.Pressed(pixelgl.KeyA) {
		Global.Ctrl.X--
	}
	if Global.Win.Pressed(pixelgl.KeyD) {
		Global.Ctrl.X++
	}
	if Global.Win.Pressed(pixelgl.KeyW) {
		Global.Ctrl.Y++
	}
	if Global.Win.Pressed(pixelgl.KeyS) {
		Global.Ctrl.Y--
	}
}

func run() {
	initScreen()

	startScene := scene.New("untitled")

	last := time.Now()
	for !Global.Win.Closed() {
		Global.DeltaTime = time.Since(last).Seconds()
		last = time.Now()

		Global.Ctrl = pixel.ZV

		handleInput()

		Global.Win.Clear(colornames.White)

		startScene.Update()
		startScene.Draw()
		Global.Win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
