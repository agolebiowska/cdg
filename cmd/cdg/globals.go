package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

var (
	pm  *pixelgl.Monitor
	win *pixelgl.Window

	worldCanvas *pixelgl.Canvas
	camPos      pixel.Vec
	cam         pixel.Matrix

	frameTick *time.Ticker
	fps       float64

	screenWidth  = 1024
	screenHeight = 768
	title        = "Companion driven Game"

	assetsPath      = "../../assets/"
	playerSheet     = assetsPath + "player.png"
	playerSheetDesc = assetsPath + "player.csv"
)
