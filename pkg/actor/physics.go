package actor

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
)

type Phys struct {
	Speed float64
	Rect  pixel.Rect

	vel pixel.Vec

	ref *Actor
}

func NewPhysics() *Phys {
	return &Phys{
		Speed: 80,
		Rect:  pixel.R(-Global.TileSize, -Global.TileSize, Global.TileSize, Global.TileSize),
	}
}

func (p *Phys) Update() {
	// apply controls
	switch {
	case Global.Ctrl.X < 0:
		p.vel.X = -p.Speed
	case Global.Ctrl.X > 0:
		p.vel.X = +p.Speed
	case Global.Ctrl.Y < 0:
		p.vel.Y = -p.Speed
	case Global.Ctrl.Y > 0:
		p.vel.Y = +p.Speed
	default:
		p.vel.X = 0
		p.vel.Y = 0
	}

	p.Rect = p.Rect.Moved(p.vel.Scaled(Global.DeltaTime))
}

func (p *Phys) SetRef(ref *Actor) {
	p.ref = ref
}

func (p *Phys) GetType() string {
	return "physics"
}
