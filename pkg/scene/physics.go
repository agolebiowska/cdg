package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Phys struct {
	Speed float64
	Rect  pixel.Rect

	vel pixel.Vec

	attack bool

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
		ph := *a.GetComponent(Physics)
		if ph == nil {
			continue
		}
		phys := ph.(*Phys)

		if phys.collide(m) {
			if a.Tag == Solid {
				return
			}
			if a.Tag == Enemy {
				if Global.Win.JustPressed(pixelgl.KeyH) {
					p.refActor.Attack(a)
				}
				return
			}
		}
	}

	p.Rect = m
}

func (p *Phys) collide(other pixel.Rect) bool {
	return p.Rect.Intersect(other) != pixel.R(0, 0, 0, 0)
}

func (p *Phys) SetRef(ref *Actor) {
	p.refActor = ref
}

func (p *Phys) GetType() componentType {
	return "physics"
}
