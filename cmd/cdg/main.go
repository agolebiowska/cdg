package main

import (
	"github.com/agolebiowska/cdg/pkg/files"
	"github.com/agolebiowska/cdg/pkg/player"
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"time"
)

func initScreen() {
	pm = pixelgl.PrimaryMonitor()
	cfg := pixelgl.WindowConfig{
		Title:   title,
		Bounds:  pixel.R(0, 0, float64(screenWidth), float64(screenHeight)),
		VSync:   true,
		Monitor: nil,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win = window
	worldCanvas = pixelgl.NewCanvas(pixel.R(-1024/2, -768/2, 1024/2, 768/2))
	camPos = pixel.ZV
}

func handleInput(ctrl *pixel.Vec) {
	if win.JustPressed(pixelgl.KeyEscape) {
		win.SetClosed(true)
	}

	// control the player with keys
	if win.Pressed(pixelgl.KeyA) {
		ctrl.X--
	}
	if win.Pressed(pixelgl.KeyD) {
		ctrl.X++
	}
	if win.Pressed(pixelgl.KeyW) {
		ctrl.Y++
	}
	if win.Pressed(pixelgl.KeyS) {
		ctrl.Y--
	}
}

func run() {
	initScreen()

	// LOAD MAP
	m, err := tilepix.ReadFile(assetsPath + "map.tmx")
	if err != nil {
		panic(err)
	}

	sheet, anims, err := files.LoadAnimationSheet(playerSheet, playerSheetDesc, 32)
	if err != nil {
		panic(err)
	}

	pp := &player.Phys{
		Speed: 80,
		Rect:  pixel.R(-32, -32, 32, 32),
	}
	pa := &player.Anim{
		Sheet: sheet,
		Anims: anims,
		Rate:  1.0 / 10,
		Dir:   +1,
	}

	imd := imdraw.New(sheet)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		// lerp the camera position towards the player
		camPos = pixel.Lerp(camPos, pp.Rect.Center(), 1-math.Pow(1.0/128, dt))
		cam = pixel.IM.Moved(camPos.Scaled(-1))
		worldCanvas.SetMatrix(cam)

		ctrl := pixel.ZV
		handleInput(&ctrl)

		pp.Update(dt, ctrl)
		pa.Update(dt, pp)

		win.Clear(colornames.White)
		worldCanvas.Clear(colornames.Black)
		worldCanvas.Draw(win, cam)

		// draw the scene to the canvas using IMDraw
		imd.Clear()
		pa.Draw(imd, pp)

		m.DrawAll(worldCanvas, color.Transparent, pixel.IM.Moved(pixel.V(-1024, -768)))

		imd.Draw(worldCanvas)

		// stretch the canvas to the window
		win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
			math.Min(
				win.Bounds().W()/worldCanvas.Bounds().W(),
				win.Bounds().H()/worldCanvas.Bounds().H(),
			),
		).Moved(win.Bounds().Center()))
		worldCanvas.Draw(win, pixel.IM.Moved(worldCanvas.Bounds().Center()))

		win.Update()
	}
}

func draw() {
	//TODO
}

func main() {
	pixelgl.Run(run)
}
