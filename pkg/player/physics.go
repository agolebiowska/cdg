package player

import "github.com/faiface/pixel"

type Phys struct {
	Speed float64
	Rect  pixel.Rect
	vel   pixel.Vec
}

func (p *Phys) Update(dt float64, ctrl pixel.Vec) {
	// apply controls
	switch {
	case ctrl.X < 0:
		p.vel.X = -p.Speed
	case ctrl.X > 0:
		p.vel.X = +p.Speed
	case ctrl.Y < 0:
		p.vel.Y = -p.Speed
	case ctrl.Y > 0:
		p.vel.Y = +p.Speed
	default:
		p.vel.X = 0
		p.vel.Y = 0
	}

	p.Rect = p.Rect.Moved(p.vel.Scaled(dt))
}
