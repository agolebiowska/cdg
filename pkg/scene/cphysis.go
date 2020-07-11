package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"log"
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

	// control the player actor with keys
	if Global.Win.Pressed(pixelgl.KeyA) {
		Global.Ctrl.X--
	}
	if Global.Win.Pressed(pixelgl.KeyD) {
		Global.Ctrl.X++
	}
	if Global.Win.Pressed(pixelgl.KeyW) {
		Global.Ctrl.Y++
	}
	if Global.Win.Pressed(pixelgl.KeyS) {
		Global.Ctrl.Y--
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

	talkingRange := p.Rect.Resized(p.Rect.Center(), pixel.V(m.W()+10, m.H()+10))
	combatRange := p.Rect.Resized(p.Rect.Center(), pixel.V(m.W()+10, m.H()+10))

	for _, a := range p.refActor.refScene.Actors {
		ph := *a.GetComponent(Physics)
		if ph == nil {
			continue
		}
		phys := ph.(*Phys)

		if phys.Rect.Intersects(talkingRange) && a.Tag == NPC {
			if Global.Win.JustPressed(pixelgl.KeyE) {
				log.Println("TALKING")
			}
		}
		if phys.Rect.Intersects(combatRange) && a.Tag == Enemy {
			if Global.Win.JustPressed(pixelgl.MouseButtonLeft) {
				c := *p.refActor.GetComponent(Combat)
				if c == nil {
					return
				}
				comb := c.(*Comb)
				comb.Attack(a)
			}
		}
		if phys.Rect.Intersects(m) && (a.Tag == Solid || a.Tag == Enemy || a.Tag == NPC) {
			return
		}
	}

	p.Rect = m
}

func (p *Phys) SetRef(ref *Actor) {
	p.refActor = ref
}

func (p *Phys) GetType() componentType {
	return "physics"
}
