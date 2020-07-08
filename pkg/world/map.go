package world

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type Map struct {
	Canvas  *pixelgl.Canvas
	Tilemap *tilepix.Map
}

func New(from string) *Map {
	m, err := tilepix.ReadFile(Global.Assets + from + ".tmx")
	if err != nil {
		panic(err)
	}

	c := pixelgl.NewCanvas(m.Bounds())

	return &Map{
		Canvas:  c,
		Tilemap: m,
	}
}

func (m *Map) Draw() {
	m.Canvas.Clear(colornames.Black)
	m.Tilemap.DrawAll(m.Canvas, color.Transparent, pixel.IM)
	// stretch the canvas to the window
	Global.Win.SetMatrix(pixel.IM.Moved(Global.Win.Bounds().Center()))
	m.Canvas.Draw(Global.Win, pixel.IM)
}
