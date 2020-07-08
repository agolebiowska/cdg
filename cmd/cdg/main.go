package main

import (
	"github.com/agolebiowska/cdg/pkg/actor"
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/agolebiowska/cdg/pkg/world"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
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

	fMap := world.New("map")
	player := actor.NewPlayer("player")

	last := time.Now()
	for !Global.Win.Closed() {
		Global.DeltaTime = time.Since(last).Seconds()
		last = time.Now()

		Global.Ctrl = pixel.ZV
		// lerp the camera position towards the actor
		Global.CamPos = pixel.Lerp(Global.CamPos, player.Phys.Rect.Center(), 1-math.Pow(1.0/128, Global.DeltaTime))
		cam := pixel.IM.Moved(Global.CamPos.Scaled(-1))
		fMap.Canvas.SetMatrix(cam)

		handleInput()
		player.Update()

		Global.Win.Clear(colornames.White)

		fMap.Draw()
		player.Draw(fMap)
		Global.Win.Update()
	}
}

func draw() {
	//TODO
}

func main() {
	pixelgl.Run(run)
}
