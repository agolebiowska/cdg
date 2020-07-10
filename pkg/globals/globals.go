package globals

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type State struct {
}

type GlobalVars struct {
	PrimaryMonitor *pixelgl.Monitor
	WindowHeight   float64
	WindowWidth    float64
	Vsync          bool
	ClearColor     color.Color
	Win            *pixelgl.Window
	Title          string
	DeltaTime      float64
	Assets         string
	TileSize       float64
	Ctrl           pixel.Vec
	CamPos         pixel.Vec
	State          *State
}

var Global = &GlobalVars{
	WindowHeight: 768,
	WindowWidth:  1024,
	Vsync:        true,
	ClearColor:   colornames.Black,
	Win:          &pixelgl.Window{},
	Title:        "Companion driven Game",
	Assets:       "../../assets/",
	TileSize:     32,
	Ctrl:         pixel.ZV,
	CamPos:       pixel.ZV,
}
