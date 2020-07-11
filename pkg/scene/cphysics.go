package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type phys struct {
	speed    float64
	rect     pixel.Rect
	vel      pixel.Vec
	attack   bool
	refActor *actor
}

func newPhysics(w, h float64) *phys {
	return &phys{
		speed: 80,
		rect:  pixel.R(-w, -h, w, h),
	}
}

func (p *phys) update() {
	if p.refActor.isPlayer == false {
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
		p.vel.X = -p.speed
	case Global.Ctrl.X > 0:
		p.vel.X = +p.speed
	case Global.Ctrl.Y < 0:
		p.vel.Y = -p.speed
	case Global.Ctrl.Y > 0:
		p.vel.Y = +p.speed
	default:
		p.vel.X = 0
		p.vel.Y = 0
	}

	m := p.rect.Moved(p.vel.Scaled(Global.DeltaTime))

	talkingRange := p.rect.Resized(p.rect.Center(), pixel.V(m.W()+10, m.H()+10))
	combatRange := p.rect.Resized(p.rect.Center(), pixel.V(m.W()+10, m.H()+10))

	for _, a := range p.refActor.refScene.actors {
		ph := *a.getComponent(Physics)
		if ph == nil {
			continue
		}
		phys := ph.(*phys)

		if phys.rect.Intersects(talkingRange) && a.tag == NPC {
			if Global.Win.JustPressed(pixelgl.KeyE) {
				d := *a.getComponent(Dialogue)
				if d == nil {
					return
				}
				dial := d.(*dial)
				dial.talk(a)
			}
		}

		if phys.rect.Intersects(combatRange) && a.tag == Enemy {
			if Global.Win.JustPressed(pixelgl.MouseButtonLeft) {
				c := *p.refActor.getComponent(Combat)
				if c == nil {
					return
				}
				comb := c.(*comb)
				comb.attack(a)
			}
		}

		if phys.rect.Intersects(m) && a.isCollidable() {
			return
		}
	}

	p.rect = m
}

func (p *phys) draw() {}

func (p *phys) setRef(ref *actor) {
	p.refActor = ref
}

func (p *phys) getType() componentType {
	return "physics"
}
