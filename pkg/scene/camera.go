package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"math"
)

type camera struct {
	pos    pixel.Vec
	follow *actor
	cam    pixel.Matrix
	zoom   float64
}

func newCamera() *camera {
	return &camera{
		pos:  pixel.ZV,
		cam:  pixel.Matrix{},
		zoom: 1,
	}
}

func (c *camera) SetPosition(x, y float64) {
	c.pos = pixel.Vec{X: x, Y: y}
}

func (c *camera) setFollow(a *actor) {
	c.follow = a
}

func (c *camera) update() {
	pos := c.pos
	if c.follow != nil {
		pos = c.follow.getPos()
	}

	pos = pixel.Lerp(c.pos, pos, 1-math.Pow(1.0/128, Global.DeltaTime))
	c.cam = pixel.IM.Moved(pos.Scaled(-1 / c.zoom))
	c.cam = c.cam.Scaled(pos, c.zoom)
	c.follow.refScene.canvas.SetMatrix(c.cam)
	c.pos = pos
}
