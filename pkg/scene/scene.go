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
		Camera:  &Camera{},
	}

	player := NewPlayer()
	scene.Add(player)

	//scene.Camera.New()
	//scene.Camera.SetFollow(player)

	// center objects same as the tilemap is centered
	center := pixel.V(-float64((scene.Tilemap.Width*scene.Tilemap.TileWidth)/2),
		-float64((scene.Tilemap.Height*scene.Tilemap.TileHeight)/2))

	// player starting position can be added in tiled as "point"
	for _, objectGroups := range scene.Tilemap.ObjectGroups {
		for _, o := range objectGroups.Objects {
			if o.Name == "player" {
				player.MoveTo(pixel.V(center.X+o.X, center.Y+o.Y))
			}
			if objectGroups.Name == "solid" {
				a := NewActor(center.X+o.X, center.Y+o.Y, o.Width, o.Height)

				a.SetTag("solid")
				scene.Actors = append(scene.Actors, a)
			}
		}
	}

	return scene
}

func (m *Scene) Add(a *Actor) {
	a.refScene = m
	m.Actors = append(m.Actors, a)
}

func (m *Scene) Draw() {
	m.Canvas.Clear(colornames.Black)

	center := pixel.V(-float64((m.Tilemap.Width*m.Tilemap.TileWidth)/2),
		-float64((m.Tilemap.Height*m.Tilemap.TileHeight)/2))
	m.Tilemap.DrawAll(m.Canvas, color.Transparent, pixel.IM.Scaled(pixel.ZV, math.Min(
		Global.Win.Bounds().W()/m.Canvas.Bounds().W(),
		Global.Win.Bounds().H()/m.Canvas.Bounds().H()),
	).Moved(center))

	for _, actor := range m.Actors {
		actor.Draw()
	}

	// DEBUG COLLIDERS
	if Global.Debug {
		for _, actor := range m.Actors {
			p := *actor.GetComponent(Physics)
			phys := p.(*Phys)
			imd := imdraw.New(nil)
			imd.Color = color.White
			imd.Push(
				phys.Rect.Norm().Vertices()[0],
				phys.Rect.Norm().Vertices()[1],
				phys.Rect.Norm().Vertices()[2],
				phys.Rect.Norm().Vertices()[3])
			imd.Polygon(1)
			imd.Draw(m.Canvas)
		}
	}

	// stretch the canvas to the window
	Global.Win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			Global.Win.Bounds().W()/m.Canvas.Bounds().W(),
			Global.Win.Bounds().H()/m.Canvas.Bounds().H(),
		),
	).Moved(Global.Win.Bounds().Center()))
	m.Canvas.Draw(Global.Win, pixel.IM.Moved(m.Canvas.Bounds().Center()))
}

func (m *Scene) Update() {
	var player *Actor
	for _, actor := range m.Actors {
		if actor.IsPlayer {
			player = actor
		}
	}
	if player == nil {
		return
	}

	//m.Camera.Update()

	// lerp the camera position towards the player
	//pos := pixel.ZV
	//pos = pixel.Lerp(pos, player.GetPos(), 1-math.Pow(1.0/128, Global.DeltaTime))
	//cam := pixel.IM.Moved(pos.Scaled(-1))
	//m.Canvas.SetMatrix(cam)

	for _, actor := range m.Actors {
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
