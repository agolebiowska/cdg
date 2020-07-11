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
	Maps           string
	Actors         string
	TileSize       float64
	Ctrl           pixel.Vec
	State          *State
	Debug          bool
}

var Global = &GlobalVars{
	//WindowHeight: 768,
	WindowHeight: 1080,
	//WindowWidth:  1024,
	WindowWidth: 1920,
	Vsync:       true,
	ClearColor:  colornames.Black,
	Win:         &pixelgl.Window{},
	Title:       "Companion driven Game",
	Assets:      "../../assets/",
	Maps:        "../../assets/maps/",
	Actors:      "../../assets/actors/",
	TileSize:    32,
	Ctrl:        pixel.ZV,
	Debug:       false,
}
