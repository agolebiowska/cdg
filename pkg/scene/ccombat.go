package scene

type comb struct {
	hp       float64
	dmg      float64
	refActor *actor
}

func newCombat(hp, dmg float64) *comb {
	return &comb{hp: hp, dmg: dmg}
}

func (l *comb) update() {
	if l.hp <= 0 {
		l.refActor.destroy()
	}
}

func (l *comb) setRef(ref *actor) {
	l.refActor = ref
}

func (l *comb) getType() componentType {
	return "combat"
}

func (l *comb) Attack(other *actor) {
	otherC := *other.getComponent(Combat)
	c := *l.refActor.getComponent(Combat)
	if otherC == nil || c == nil {
		return
	}
	otherCombat := otherC.(*comb)
	combat := c.(*comb)
	otherCombat.hp -= combat.dmg
}
