package scene

import (
	"fmt"
	"github.com/agolebiowska/cdg/pkg/files"
	. "github.com/agolebiowska/cdg/pkg/globals"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"strings"
)

type link struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type passage struct {
	Text  string  `json:"text"`
	Links []*link `json:"links"`
	Name  string  `json:"name"`
	Pid   int     `json:"pid,string"`
}

type dial struct {
	Passages  []*passage `json:"passages"`
	StartNode int        `json:"startnode,string"`
	//
	passages   map[int]*passage
	counter    int
	on         bool
	refActor   *actor
	speakingTo *actor
}

func newDialogue(from string) *dial {
	var dl *dial
	files.LoadJSON(Global.Dialogues+from+".json", &dl)
	dl.passages = map[int]*passage{}
	dl.counter = dl.StartNode

	for _, p := range dl.Passages {
		dl.passages[p.Pid] = p
	}

	return dl
}

func (d *dial) update() {
}

func (d *dial) draw() {
	if d.on {
		p := d.passages[d.counter]

		basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
		basicTxt := text.New(pixel.V(d.refActor.getPos().X-100, d.refActor.getPos().Y+40), basicAtlas)
		answers := text.New(pixel.V(d.speakingTo.getPos().X, d.speakingTo.getPos().Y+30), basicAtlas)

		fmt.Fprintln(basicTxt, p.Name)
		for i, l := range p.Links {
			fmt.Fprintln(answers, fmt.Sprintf("%d. %s", i+1, strings.Split(l.Name, "|")[0]))
		}

		basicTxt.Draw(d.refActor.refScene.canvas, pixel.IM)
		answers.Draw(d.refActor.refScene.canvas, pixel.IM)
	}
}

func (d *dial) setRef(ref *actor) {
	d.refActor = ref
}

func (d *dial) getType() componentType {
	return "dialogue"
}

func (d *dial) talk(to *actor) {
	d.speakingTo = to
	if len(d.passages) <= d.counter {
		d.on = false
		d.counter = 0
		return
	}

	d.on = true
	d.counter++
}
