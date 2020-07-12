package scene

import (
	"fmt"
	"github.com/faiface/pixel"
	"reflect"
)

type actorType string

var (
	Player actorType = "player"
	Enemy  actorType = "enemy"
	NPC    actorType = "npc"
	Solid  actorType = "solid"
)

type actor struct {
	tag        actorType
	isPlayer   bool
	components []component
	refScene   *scene
}

func newPlayer() *actor {
	actor := &actor{
		tag:        Player,
		isPlayer:   true,
		components: []component{},
	}

	actor.addComponent(newPhysics(10, 15))
	actor.addComponent(newAnim("player"))
	actor.addComponent(newCombat(100, 10))

	return actor
}

func newActor(x, y, w, h float64) *actor {
	actor := &actor{
		isPlayer: false,
	}

	actor.addComponent(newPhysics(w, h))
	actor.moveTo(pixel.V((x*2)+w, (y*2)+h))

	return actor
}

func (a *actor) setTag(tag actorType) {
	a.tag = tag
}

func (a *actor) addComponent(c component) {
	for _, existing := range a.components {
		if reflect.TypeOf(c) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type: %v",
				reflect.TypeOf(c)))
		}
	}
	c.setRef(a)
	a.components = append(a.components, c)
}

func (a *actor) getComponent(t componentType) *component {
	for _, c := range a.components {
		if c.getType() == t {
			return &c
		}
	}
	return nil
}

func (a *actor) update() {
	for _, component := range a.components {
		component.update()
	}
}

func (a *actor) draw() {
	var anm *anim
	var phs *phys
	var dlg *dial
	for _, c := range a.components {
		if c.getType() == Animation {
			anm = c.(*anim)
		}
		if c.getType() == Physics {
			phs = c.(*phys)
		}
		if c.getType() == Dialogue {
			dlg = c.(*dial)
		}
	}

	if dlg != nil {
		dlg.draw()
	}

	if anm == nil || phs == nil {
		return
	}

	if anm.sprite == nil {
		anm.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	// draw the correct frame with the correct position and direction
	anm.sprite.Set(anm.sheet, anm.frame)
	anm.sprite.Draw(a.refScene.canvas, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			anm.sprite.Frame().W()/32,
			anm.sprite.Frame().H()/32,
		)).
		ScaledXY(pixel.ZV, pixel.V(anm.dir, 1)).Moved(phs.rect.Center()),
	)
}

func (a *actor) moveTo(vec pixel.Vec) {
	p := *a.getComponent(Physics)
	if p == nil {
		return
	}
	phs := p.(*phys)

	phs.rect = phs.rect.Moved(vec)
}

func (a *actor) getPos() pixel.Vec {
	p := *a.getComponent(Physics)
	if p == nil {
		return pixel.V(0, 0)
	}
	phs := p.(*phys)

	return phs.rect.Center()
}

func (a *actor) destroy() {
	a.components = []component{}
	for i, ac := range a.refScene.actors {
		if len(ac.components) <= 0 {
			a.refScene.actors = append(a.refScene.actors[:i], a.refScene.actors[i+1:]...)
		}
	}
}

func (a *actor) isCollidable() bool {
	return a.tag == Solid ||
		a.tag == Enemy ||
		a.tag == NPC
}
