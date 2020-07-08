package player

import (
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

func (pa *Anim) Update(dt float64, phys *Phys) {
	pa.counter += dt

	// determine the new animation state
	var newState animState
	switch {
	case phys.vel.Len() == 0:
		newState = idle
	case phys.vel.Len() > 0:
		newState = running
	}

	// reset the time counter if the state changed
	if pa.state != newState {
		pa.state = newState
		pa.counter = 0
	}

	// determine the correct animation frame
	switch pa.state {
	case idle:
		//pa.frame = pa.Anims["Front"][0]
		i := int(math.Floor(pa.counter / pa.Rate))
		pa.frame = pa.Anims["Front"][i%len(pa.Anims["Front"])]
	case running:
		i := int(math.Floor(pa.counter / pa.Rate))
		pa.frame = pa.Anims["Run"][i%len(pa.Anims["Run"])]
	}

	// set the facing direction of the gopher
	if phys.vel.X != 0 {
		if phys.vel.X > 0 {
			pa.Dir = +1
		} else {
			pa.Dir = -1
		}
	}
}

func (pa *Anim) Draw(t pixel.Target, phys *Phys) {
	if pa.sprite == nil {
		pa.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	// draw the correct frame with the correct position and direction
	pa.sprite.Set(pa.Sheet, pa.frame)
	pa.sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			phys.Rect.W()/pa.sprite.Frame().W(),
			phys.Rect.H()/pa.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(pa.Dir, 1)).
		Moved(phys.Rect.Center()),
	)
}
