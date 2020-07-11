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

type Scene struct {
	Canvas  *pixelgl.Canvas
	Tilemap *tilepix.Map
	Actors  []*Actor
	Camera  *Camera
}

func New(from string) *Scene {
	m, err := tilepix.ReadFile(Global.Assets + from + ".tmx")
	if err != nil {
		panic(err)
	}

	c := pixelgl.NewCanvas(pixel.R(
		-Global.WindowWidth/4,
		-Global.WindowHeight/4,
		Global.WindowWidth/4,
		Global.WindowHeight/4))

	scene := &Scene{
		Canvas:  c,
		Tilemap: m,
		Actors:  []*Actor{},
		Camera:  NewCamera(),
	}

	player := NewPlayer()
	scene.Add(player)
	scene.Camera.SetFollow(player)

	// center objects same as the tilemap is centered
	center := pixel.V(-float64((scene.Tilemap.Width*scene.Tilemap.TileWidth)/2),
		-float64((scene.Tilemap.Height*scene.Tilemap.TileHeight)/2))

	// player starting position can be added in tiled as "point"
	for _, objectGroups := range scene.Tilemap.ObjectGroups {
		for _, o := range objectGroups.Objects {
			if o.Name == "player" {
				player.MoveTo(pixel.V(center.X+o.X, center.Y+o.Y))
			}

			if objectGroups.Name == "enemy" {
				a := NewActor(center.X+o.X, center.Y+o.Y, 10, 15)
				a.AddComponent(NewAnim("enemy"))
				a.AddComponent(NewCombat(100, 10))
				a.SetTag(Enemy)
				scene.Add(a)
			}

			if objectGroups.Name == "solid" {
				a := NewActor(center.X+o.X, center.Y+o.Y, o.Width, o.Height)
				a.SetTag(Solid)
				scene.Add(a)
			}
		}
	}

	return scene
}

func (s *Scene) Add(a *Actor) {
	a.refScene = s
	s.Actors = append(s.Actors, a)
}

func (s *Scene) Draw() {
	s.Canvas.Clear(colornames.Black)

	center := pixel.V(-float64((s.Tilemap.Width*s.Tilemap.TileWidth)/2),
		-float64((s.Tilemap.Height*s.Tilemap.TileHeight)/2))

	s.Tilemap.DrawAll(s.Canvas, color.Transparent, pixel.IM.Scaled(pixel.ZV, math.Min(
		Global.Win.Bounds().W()/s.Canvas.Bounds().W(),
		Global.Win.Bounds().H()/s.Canvas.Bounds().H()),
	).Moved(center))

	for _, actor := range s.Actors {
		actor.Draw()
	}

	// DEBUG COLLIDERS
	if Global.Debug {
		for _, actor := range s.Actors {
			p := *actor.GetComponent(Physics)
			if p != nil {
				phys := p.(*Phys)
				imd := imdraw.New(nil)
				imd.Color = color.White
				imd.Push(
					phys.Rect.Norm().Vertices()[0],
					phys.Rect.Norm().Vertices()[1],
					phys.Rect.Norm().Vertices()[2],
					phys.Rect.Norm().Vertices()[3])
				imd.Polygon(1)
				imd.Draw(s.Canvas)
			}
		}
	}

	// stretch the canvas to the window
	Global.Win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			Global.Win.Bounds().W()/s.Canvas.Bounds().W(),
			Global.Win.Bounds().H()/s.Canvas.Bounds().H(),
		),
	).Moved(Global.Win.Bounds().Center()))
	s.Canvas.Draw(Global.Win, pixel.IM.Moved(s.Canvas.Bounds().Center()))
}

func (s *Scene) Update() {
	s.Camera.Update()

	for _, actor := range s.Actors {
		actor.Update()
	}
}

func (s *Scene) GetPlayer() *Actor {
	for _, a := range s.Actors {
		if a.IsPlayer {
			return a
		}
	}
	return nil
}
