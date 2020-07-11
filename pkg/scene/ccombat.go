package scene

type comb struct {
	hp       float64
	dmg      float64
	refActor *actor
}

func newCombat(hp, dmg float64) *comb {
	return &comb{hp: hp, dmg: dmg}
}

func (c *comb) update() {
	if c.hp <= 0 {
		c.refActor.destroy()
	}
}

func (c *comb) draw() {}

func (c *comb) setRef(ref *actor) {
	c.refActor = ref
}

func (c *comb) getType() componentType {
	return "combat"
}

func (c *comb) attack(other *actor) {
	otherC := *other.getComponent(Combat)
	cmb := *c.refActor.getComponent(Combat)
	if otherC == nil || c == nil {
		return
	}
	otherCombat := otherC.(*comb)
	combat := cmb.(*comb)
	otherCombat.hp -= combat.dmg
}
