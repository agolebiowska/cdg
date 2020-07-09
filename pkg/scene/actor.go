package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
)

type Actor struct {
	Tag        string
	IsPlayer   bool
	Components []Component
}

func NewPlayer() *Actor {
	actor := &Actor{
		Tag:        "player",
		IsPlayer:   true,
		Components: []Component{},
	}

	actor.AddComponent(NewPhysics())
	actor.AddComponent(NewAnim("player"))

	return actor
}

func NewActor(vec pixel.Vec) *Actor {
	actor := &Actor{
		IsPlayer: false,
	}

	actor.AddComponent(NewPhysics())
	actor.MoveTo(vec)

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
		if c.GetType() == "animation" {
			anim = c.(*Anim)
		}
		if c.GetType() == "physics" {
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
	var phys *Phys
	for _, c := range a.Components {
		if c.GetType() == "physics" {
			phys = c.(*Phys)
		}
	}
	if phys == nil {
		return
	}
	phys.Rect = phys.Rect.Moved(vec)
}

func (a *Actor) GetPos() pixel.Vec {
	var phys *Phys
	for _, c := range a.Components {
		if c.GetType() == "physics" {
			phys = c.(*Phys)
		}
	}
	if phys == nil {
		return pixel.V(0, 0)
	}
	return phys.Rect.Center()
}
