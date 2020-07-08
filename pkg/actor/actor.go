package actor

import (
	"github.com/agolebiowska/cdg/pkg/files"
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/agolebiowska/cdg/pkg/world"
	"github.com/faiface/pixel"
)

type Actor struct {
	IsPlayer bool
	Phys     *Phys
	Anim     *Anim
}

func NewPlayer(from string) *Actor {
	sheet, anims, err := files.LoadAnimationSheet(
		Global.Assets+from+".png",
		Global.Assets+from+".csv",
		Global.TileSize,
	)
	if err != nil {
		panic(err)
	}

	return &Actor{
		IsPlayer: true,
		Phys: &Phys{
			Speed: 80,
			Rect:  pixel.R(-Global.TileSize, -Global.TileSize, Global.TileSize, Global.TileSize),
		},
		Anim: &Anim{
			Sheet: sheet,
			Anims: anims,
			Rate:  1.0 / 10,
			Dir:   +1,
		},
	}
}

func (a *Actor) Update() {
	if a.IsPlayer {
		a.Phys.update()
	}
	a.Anim.update(a.Phys)
}

func (a *Actor) Draw(where *world.Map) {
	if a.Anim.sprite == nil {
		a.Anim.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	// draw the correct frame with the correct position and direction
	a.Anim.sprite.Set(a.Anim.Sheet, a.Anim.frame)
	a.Anim.sprite.Draw(Global.Win, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			a.Phys.Rect.W()/a.Anim.sprite.Frame().W(),
			a.Phys.Rect.H()/a.Anim.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(a.Anim.Dir, 1)),
	)
}
