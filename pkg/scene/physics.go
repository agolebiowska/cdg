package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
)

type Phys struct {
	Speed float64
	Rect  pixel.Rect

	vel pixel.Vec

	refActor *Actor
}

func NewPhysics(w, h float64) *Phys {
	return &Phys{
		Speed: 80,
		Rect:  pixel.R(-w, -h, w, h),
	}
}

func (p *Phys) Update() {
	if p.refActor.IsPlayer == false {
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

	for _, a := range p.refActor.refScene.Actors {
		if a.Tag == "solid" {
			p := *a.GetComponent("physics")
			if p == nil {
				continue
			}
			phys := p.(*Phys)
			if phys.Rect.Intersect(m) != pixel.R(0, 0, 0, 0) {
				return
			}
		}
	}

	p.Rect = m
}

func (p *Phys) SetRef(ref *Actor) {
	p.refActor = ref
}

func (p *Phys) GetType() string {
	return "physics"
}
