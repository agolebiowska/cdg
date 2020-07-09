package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"log"
	"math"
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
	if p.ref.IsPlayer == false {
		return
	}

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

	m := p.Rect.Moved(p.vel.Scaled(Global.DeltaTime))
	X := m.Center().X
	Y := m.Center().Y
	log.Println(Global.State.MapData)
	log.Println(Global.State.MapData[pixel.V(math.Round(X), math.Round(Y))])
	if Global.State.MapData[pixel.V(math.Round(X), math.Round(Y))] == "solid" {
		log.Println(pixel.V(math.Round(X), math.Round(Y)))
		return
	}

	p.Rect = m
}

func (p *Phys) SetRef(ref *Actor) {
	p.ref = ref
}

func (p *Phys) GetType() string {
	return "physics"
}
