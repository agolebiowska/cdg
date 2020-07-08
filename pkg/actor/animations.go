package actor

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"math"
)

type animState int

const (
	idle animState = iota
	running
)

type Anim struct {
	Sheet pixel.Picture
	Anims map[string][]pixel.Rect
	Rate  float64

	state   animState
	counter float64
	Dir     float64

	frame pixel.Rect

	sprite *pixel.Sprite
}

func (a *Anim) update(phys *Phys) {
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
		a.frame = a.Anims["Front"][i%len(a.Anims["Front"])]
	case running:
		i := int(math.Floor(a.counter / a.Rate))
		a.frame = a.Anims["Run"][i%len(a.Anims["Run"])]
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
