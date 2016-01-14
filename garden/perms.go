package garden

type permanent interface {
}

type minionPerm struct {
	attack               int
	health               int
	hasSummoningSickness bool
}

func newMinionPerm(mc *MinionCard) *minionPerm {
	m := new(minionPerm)
	m.attack = mc.Attack
	m.health = mc.Health
	m.hasSummoningSickness = true
	return m
}
