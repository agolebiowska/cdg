package scene

import (
	"github.com/faiface/pixel"
)

type actorType string

var (
	Player actorType = "player"
	Enemy  actorType = "enemy"
	NPC    actorType = "npc"
	Solid  actorType = "solid"
)

type Actor struct {
	Tag        actorType
	IsPlayer   bool
	Components []Component

	refScene *Scene
}

func NewPlayer() *Actor {
	actor := &Actor{
		Tag:        Player,
		IsPlayer:   true,
		Components: []Component{},
	}

	actor.AddComponent(NewPhysics(10, 15))
	actor.AddComponent(NewAnim("player"))
	actor.AddComponent(NewCombat(100, 10))

	return actor
}

func NewActor(x, y, w, h float64) *Actor {
	actor := &Actor{
		IsPlayer: false,
	}

	actor.AddComponent(NewPhysics(w, h))
	actor.MoveTo(pixel.V((x*2)+w, (y*2)+h))

	return actor
}

func (a *Actor) SetTag(tag actorType) {
	a.Tag = tag
}

func (a *Actor) AddComponent(c Component) {
	c.SetRef(a)
	a.Components = append(a.Components, c)
}

func (a *Actor) Update() {
	for _, component := range a.Components {
		component.Update()
	}
}

func (a *Actor) Draw() {
	var anim *Anim
	var phys *Phys
	for _, c := range a.Components {
		if c.GetType() == Animation {
			anim = c.(*Anim)
		}
		if c.GetType() == Physics {
			phys = c.(*Phys)
		}
	}
	if anim == nil || phys == nil {
		return
	}

	if anim.Sprite == nil {
		anim.Sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	// draw the correct frame with the correct position and direction
	anim.Sprite.Set(anim.Sheet, anim.Frame)
	anim.Sprite.Draw(a.refScene.Canvas, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			anim.Sprite.Frame().W()/32,
			anim.Sprite.Frame().H()/32,
			//phys.Rect.W()/anim.Sprite.Frame().W(),
			//phys.Rect.H()/anim.Sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(anim.Dir, 1)).Moved(phys.Rect.Center()),
	)
}

func (a *Actor) MoveTo(vec pixel.Vec) {
	p := *a.GetComponent(Physics)
	if p == nil {
		return
	}
	phys := p.(*Phys)

	phys.Rect = phys.Rect.Moved(vec)
}

func (a *Actor) GetPos() pixel.Vec {
	p := *a.GetComponent(Physics)
	if p == nil {
		return pixel.V(0, 0)
	}
	phys := p.(*Phys)

	return phys.Rect.Center()
}

func (a *Actor) GetComponent(t componentType) *Component {
	for _, c := range a.Components {
		if c.GetType() == t {
			return &c
		}
	}
	return nil
}

func (a *Actor) Destroy() {
	a.Components = []Component{}
	for i, ac := range a.refScene.Actors {
		if len(ac.Components) <= 0 {
			a.refScene.Actors = append(a.refScene.Actors[:i], a.refScene.Actors[i+1:]...)
		}
	}
}

func (a *Actor) isCollidable() bool {
	return a.Tag == Solid ||
		a.Tag == Enemy ||
		a.Tag == NPC
}
