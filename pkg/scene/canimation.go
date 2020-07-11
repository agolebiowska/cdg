package scene

import (
	"github.com/agolebiowska/cdg/pkg/files"
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"math"
)

type animState int

const (
	idle    animState = iota
	running animState = iota
)

type Anim struct {
	Sheet pixel.Picture
	Anims map[string][]pixel.Rect
	Rate  float64

	state   animState
	counter float64
	Dir     float64

	Frame  pixel.Rect
	Sprite *pixel.Sprite

	refActor *Actor
}

func NewAnim(from string) *Anim {
	sheet, anims, err := files.LoadAnimationSheet(
		Global.Assets+from+".png",
		Global.Assets+from+".csv",
		Global.TileSize,
	)
	if err != nil {
		panic(err)
	}

	return &Anim{
		Sheet: sheet,
		Anims: anims,
		Rate:  1.0 / 10,
		Dir:   +1,
	}
}

func (a *Anim) SetRef(ref *Actor) {
	a.refActor = ref
}

func (a *Anim) Update() {
	p := *a.refActor.GetComponent(Physics)
	if p == nil {
		return
	}
	phys := p.(*Phys)

	a.counter += Global.DeltaTime

	// determine the new animation state
	var newState animState
	switch {
	case phys.vel.Len() == 0:
		newState = idle
	case phys.vel.Len() > 0:
		newState = running
	}

	// reset the time counter if the state changed
	if a.state != newState {
		a.state = newState
		a.counter = 0
	}

	// determine the correct animation frame
	switch a.state {
	case idle:
		//a.frame = a.Anims["Front"][0]
		i := int(math.Floor(a.counter / a.Rate))
		a.Frame = a.Anims["Front"][i%len(a.Anims["Front"])]
	case running:
		i := int(math.Floor(a.counter / a.Rate))
		a.Frame = a.Anims["Run"][i%len(a.Anims["Run"])]
	}

	// set the facing direction of the actor
	if phys.vel.X != 0 {
		if phys.vel.X > 0 {
			a.Dir = +1
		} else {
			a.Dir = -1
		}
	}
}

func (a *Anim) GetType() componentType {
	return "animation"
}
