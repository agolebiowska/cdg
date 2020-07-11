package main

import (
	"flag"
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/agolebiowska/cdg/pkg/scene"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
	"time"
)

var (
	debug = flag.Bool("debug", false, "Debug mode on/off")
)

func main() {
	flag.Parse()
	Global.Debug = *debug

	pixelgl.Run(run)
}

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

	startScene := scene.New("map")

	camPos := pixel.ZV

	last := time.Now()
	for !Global.Win.Closed() {
		Global.DeltaTime = time.Since(last).Seconds()
		last = time.Now()

		// lerp the camera position towards the gopher
		camPos = pixel.Lerp(camPos, startScene.GetPlayer().GetPos(), 1-math.Pow(1.0/128, Global.DeltaTime))
		cam := pixel.IM.Moved(camPos.Scaled(-1))
		startScene.Canvas.SetMatrix(cam)

		Global.Ctrl = pixel.ZV

		handleInput()

		Global.Win.Clear(colornames.Black)

		startScene.Update()
		startScene.Draw()
		Global.Win.Update()
	}
}
