package garden

import "fmt"

type permanent interface {
	isActionable() bool
	setActionable(b bool)
	act(gs *gameState)
}

type permanents []permanent

func (ps permanents) findActionablePerms() permanents {
	matches := make([]permanent, 0)
	for _, p := range ps {
		if p.isActionable() {
			matches = append(matches, p)
		}
	}
	return matches
}

type minionPerm struct {
	attack        int
	health        int
	_isActionable bool
}

func newMinionPerm(mc *MinionCard) *minionPerm {
	m := new(minionPerm)
	m.attack = mc.Attack
	m.health = mc.Health
	m._isActionable = false
	return m
}

func (m *minionPerm) isActionable() bool {
	return m._isActionable
}

func (m *minionPerm) setActionable(b bool) {
	m._isActionable = b
}

func (m *minionPerm) String() string {
	return fmt.Sprintf("minion<%v/%v>", m.attack, m.health)
}

func (m *minionPerm) act(gs *gameState) {
	m.setActionable(false)

	fmt.Fprintf(gs.logger, "Attack from %v to %v for %v health\n", m, gs.pe, m.attack)

	// attack enemy
	gs.pe.health -= m.attack
}
