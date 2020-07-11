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

type anim struct {
	sheet    pixel.Picture
	anims    map[string][]pixel.Rect
	rate     float64
	state    animState
	counter  float64
	dir      float64
	frame    pixel.Rect
	sprite   *pixel.Sprite
	refActor *actor
}

func newAnim(from string) *anim {
	sheet, anims, err := files.LoadAnimationSheet(
		Global.Actors+from+".png",
		Global.Actors+from+".csv",
		Global.TileSize,
	)
	if err != nil {
		panic(err)
	}

	return &anim{
		sheet: sheet,
		anims: anims,
		rate:  1.0 / 10,
		dir:   +1,
	}
}

func (a *anim) setRef(ref *actor) {
	a.refActor = ref
}

func (a *anim) update() {
	p := *a.refActor.getComponent(Physics)
	if p == nil {
		return
	}
	phys := p.(*phys)

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
		//a.frame = a.anims["Front"][0]
		i := int(math.Floor(a.counter / a.rate))
		a.frame = a.anims["Front"][i%len(a.anims["Front"])]
	case running:
		i := int(math.Floor(a.counter / a.rate))
		a.frame = a.anims["Run"][i%len(a.anims["Run"])]
	}

	// set the facing direction of the actor
	if phys.vel.X != 0 {
		if phys.vel.X > 0 {
			a.dir = +1
		} else {
			a.dir = -1
		}
	}
}

func (a *anim) draw() {}

func (a *anim) getType() componentType {
	return "animation"
}
