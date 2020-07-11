package scene

import "log"

type Dial struct {
	Sentences []string
	counter   int

	refActor *Actor
}

func NewDialogue(s []string) *Dial {
	return &Dial{
		Sentences: s,
		counter:   0,
	}
}

func (d *Dial) Update() {

}

func (d *Dial) SetRef(ref *Actor) {
	d.refActor = ref
}

func (d *Dial) GetType() componentType {
	return "dialogue"
}

func (d *Dial) Talk(other *Actor) {
	//otherD := *other.GetComponent(Dialogue)
	//if otherD == nil {
	//	return
	//}
	//otherDial := otherD.(*Comb)
	if len(d.Sentences) <= d.counter {
		d.counter = 0
	}
	log.Println(d.Sentences[d.counter])
	d.counter++
}
