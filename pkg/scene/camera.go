package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"math"
)

type Camera struct {
	Pos    pixel.Vec
	Follow *Actor
	Cam    pixel.Matrix
	Zoom   float64
}

func NewCamera() *Camera {
	return &Camera{
		Pos:  pixel.ZV,
		Cam:  pixel.Matrix{},
		Zoom: 1,
	}
}

func (c *Camera) SetPosition(x, y float64) {
	c.Pos = pixel.Vec{X: x, Y: y}
}

func (c *Camera) SetFollow(a *Actor) {
	c.Follow = a
}

func (c *Camera) Update() {
	pos := c.Pos
	if c.Follow != nil {
		pos = c.Follow.GetPos()
	}

	pos = pixel.Lerp(c.Pos, pos, 1-math.Pow(1.0/128, Global.DeltaTime))
	c.Cam = pixel.IM.Moved(pos.Scaled(-1 / c.Zoom))
	c.Cam = c.Cam.Scaled(pos, c.Zoom)
	c.Follow.refScene.Canvas.SetMatrix(c.Cam)
	c.Pos = pos
}
