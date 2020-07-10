package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
)

type Actor struct {
	Tag        string
	IsPlayer   bool
	Components []Component

	refScene *Scene
}

func NewPlayer() *Actor {
	actor := &Actor{
		Tag:        "player",
		IsPlayer:   true,
		Components: []Component{},
	}

	actor.AddComponent(NewPhysics(16, 16))
	actor.AddComponent(NewAnim("player"))

	return actor
}

func NewActor(x, y, w, h float64) *Actor {
	actor := &Actor{
		IsPlayer: false,
	}

	actor.AddComponent(NewPhysics(w, h))
	actor.MoveTo(pixel.V(x+w, y+h))

	return actor
}

func (a *Actor) SetTag(tag string) {
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
	if anim == nil {
		return
	}

	if anim.Sprite == nil {
		anim.Sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	// draw the correct frame with the correct position and direction
	anim.Sprite.Set(anim.Sheet, anim.Frame)
	anim.Sprite.Draw(Global.Win, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			phys.Rect.W()/anim.Sprite.Frame().W(),
			phys.Rect.H()/anim.Sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(anim.Dir, 1)),
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
