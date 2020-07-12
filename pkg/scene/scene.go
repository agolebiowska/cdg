package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type scene struct {
	canvas  *pixelgl.Canvas
	tilemap *tilepix.Map
	actors  []*actor
	camera  *camera
}

func New(from string) *scene {
	m, err := tilepix.ReadFile(Global.Maps + from + ".tmx")
	if err != nil {
		panic(err)
	}

	c := pixelgl.NewCanvas(pixel.R(
		-Global.WindowWidth/4,
		-Global.WindowHeight/4,
		Global.WindowWidth/4,
		Global.WindowHeight/4))

	scene := &scene{
		canvas:  c,
		tilemap: m,
		actors:  []*actor{},
		camera:  newCamera(),
	}

	player := newPlayer()
	scene.add(player)
	scene.camera.setFollow(player)

	center := scene.getMapCenter()

	for _, objectGroups := range scene.tilemap.ObjectGroups {
		for _, o := range objectGroups.Objects {
			switch o.Type {
			case "player":
				player.moveTo(pixel.V(center.X+o.X, center.Y+o.Y))

			case "enemy":
				a := newActor(center.X+o.X, center.Y+o.Y, 10, 15)
				a.addComponent(newAnim(o.Name))
				a.addComponent(newCombat(100, 10))
				a.setTag(Enemy)
				scene.add(a)

			case "npc":
				a := newActor(center.X+o.X, center.Y+o.Y, 16, 16)
				a.addComponent(newAnim(o.Name))
				a.addComponent(newDialogue("npc0"))
				a.setTag(NPC)
				scene.add(a)
			}

			if objectGroups.Name == "solid" {
				a := newActor(center.X+o.X, center.Y+o.Y, o.Width, o.Height)
				a.setTag(Solid)
				scene.add(a)
			}
		}
	}

	return scene
}

func (s *scene) Draw() {
	s.canvas.Clear(colornames.Black)

	s.tilemap.DrawAll(s.canvas, color.Transparent, pixel.IM.Scaled(pixel.ZV, math.Min(
		Global.Win.Bounds().W()/s.canvas.Bounds().W(),
		Global.Win.Bounds().H()/s.canvas.Bounds().H()),
	).Moved(s.getMapCenter()))

	for _, actor := range s.actors {
		actor.draw()
	}

	// DEBUG COLLIDERS
	if Global.Debug {
		for _, actor := range s.actors {
			p := *actor.getComponent(Physics)
			if p != nil {
				phys := p.(*phys)
				imd := imdraw.New(nil)
				imd.Color = color.White
				imd.Push(
					phys.rect.Norm().Vertices()[0],
					phys.rect.Norm().Vertices()[1],
					phys.rect.Norm().Vertices()[2],
					phys.rect.Norm().Vertices()[3])
				imd.Polygon(1)
				imd.Draw(s.canvas)
			}
		}
	}

	// stretch the canvas to the window
	Global.Win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			Global.Win.Bounds().W()/s.canvas.Bounds().W(),
			Global.Win.Bounds().H()/s.canvas.Bounds().H(),
		),
	).Moved(Global.Win.Bounds().Center()))
	s.canvas.Draw(Global.Win, pixel.IM.Moved(s.canvas.Bounds().Center()))
}

func (s *scene) Update() {
	s.camera.update()
	for _, actor := range s.actors {
		actor.update()
	}
}

func (s *scene) add(a *actor) {
	a.refScene = s
	s.actors = append(s.actors, a)
}

func (s *scene) getPlayer() *actor {
	for _, a := range s.actors {
		if a.isPlayer {
			return a
		}
	}
	return nil
}

// get tilemap center point
func (s *scene) getMapCenter() pixel.Vec {
	return pixel.V(
		-float64((s.tilemap.Width*s.tilemap.TileWidth)/2),
		-float64((s.tilemap.Height*s.tilemap.TileHeight)/2))
}
