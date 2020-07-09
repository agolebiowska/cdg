package scene

import (
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type Scene struct {
	Canvas  *pixelgl.Canvas
	Tilemap *tilepix.Map
	Actors  []*Actor
}

func New(from string) *Scene {
	m, err := tilepix.ReadFile(Global.Assets + from + ".tmx")
	if err != nil {
		panic(err)
	}

	c := pixelgl.NewCanvas(pixel.R(-1024/2, -768/2, 1024/2, 768/2))

	scene := &Scene{
		Canvas:  c,
		Tilemap: m,
		Actors:  []*Actor{},
	}

	player := NewPlayer()
	scene.Add(player)

	//player starting position can be added in tiled as "point"
	for _, objectGroups := range scene.Tilemap.ObjectGroups {
		for _, o := range objectGroups.Objects {
			if o.Name == "player" {
				player.MoveTo(pixel.V(o.X, o.Y))
			}
			if objectGroups.Name == "solid" {
				a := NewActor(pixel.V(o.X, o.Y))
				a.SetTag("solid")
				scene.Actors = append(scene.Actors, a)
			}
		}
	}

	for _, actor := range scene.Actors {
		if actor.Tag == "solid" {
			Global.State.MapData[pixel.V(math.Round(actor.GetPos().X), math.Round(actor.GetPos().Y))] = actor.Tag
		}
	}

	return scene
}

func (m *Scene) Add(a *Actor) {
	m.Actors = append(m.Actors, a)
}

func (m *Scene) Draw() {
	m.Canvas.Clear(colornames.Black)

	m.Tilemap.DrawAll(m.Canvas, color.Transparent, pixel.IM)
	// stretch the canvas to the window
	Global.Win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			Global.Win.Bounds().W()/m.Canvas.Bounds().W(),
			Global.Win.Bounds().H()/m.Canvas.Bounds().H(),
		),
	).Moved(Global.Win.Bounds().Center()))
	m.Canvas.Draw(Global.Win, pixel.IM.Moved(m.Canvas.Bounds().Center()))

	for _, actor := range m.Actors {
		actor.Draw()
	}
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
	// lerp the camera position towards the player
	Global.CamPos = pixel.Lerp(Global.CamPos, player.GetPos(), 1-math.Pow(1.0/128, Global.DeltaTime))
	cam := pixel.IM.Moved(Global.CamPos.Scaled(-1))
	m.Canvas.SetMatrix(cam)

	for _, actor := range m.Actors {
		actor.Update()
	}
}
