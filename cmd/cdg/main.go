package main

import (
	"github.com/agolebiowska/cdg/files"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	assetsPath = "../../assets"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Companion driven Game",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	playerPic := files.LoadPicture(assetsPath + "player.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(playerPic, playerPic.Bounds())
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
