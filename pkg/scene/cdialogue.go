package scene

import "log"

type dial struct {
	sentences []string
	counter   int
	refActor  *actor
}

func newDialogue(s []string) *dial {
	return &dial{
		sentences: s,
		counter:   0,
	}
}

func (d *dial) update() {
}

func (d *dial) setRef(ref *actor) {
	d.refActor = ref
}

func (d *dial) getType() componentType {
	return "dialogue"
}

func (d *dial) Talk(other *actor) {
	//otherD := *other.getComponent(Dialogue)
	//if otherD == nil {
	//	return
	//}
	//otherDial := otherD.(*Comb)
	if len(d.sentences) <= d.counter {
		d.counter = 0
	}
	log.Println(d.sentences[d.counter])
	d.counter++
}
