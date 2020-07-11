package scene

type Comb struct {
	HP  float64
	Dmg float64

	refActor *Actor
}

func NewCombat(HP, Dmg float64) *Comb {
	return &Comb{HP: HP, Dmg: Dmg}
}

func (l *Comb) Update() {
	if l.HP <= 0 {
		l.refActor.Destroy()
	}
}

func (l *Comb) SetRef(ref *Actor) {
	l.refActor = ref
}

func (l *Comb) GetType() componentType {
	return "combat"
}

func (l *Comb) Attack(other *Actor) {
	otherC := *other.GetComponent(Combat)
	c := *l.refActor.GetComponent(Combat)
	if otherC == nil || c == nil {
		return
	}
	otherCombat := otherC.(*Comb)
	combat := c.(*Comb)

	otherCombat.HP -= combat.Dmg
}
