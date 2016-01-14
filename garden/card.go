package garden

import (
	"fmt"
	"math"
	"math/rand"
)

type Card interface {
	Cost() int
	play(gs *gameState)
}

func (cs Cards) withinManaRange(low int, high int) Cards {
	matches := make([]Card, 0)
	for _, c := range cs {
		if low <= c.Cost() && c.Cost() <= high {
			matches = append(matches, c)
		}
	}
	return matches
}

type Cards []Card

func (cs Cards) Shuffle() {
	// The Fisher-Yates shuffling algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	for i := range cs {
		j := rand.Intn(i + 1)
		cs[i], cs[j] = cs[j], cs[i]
	}
}

func (cs Cards) RemoveCard(other Card) {
	for i, c := range cs {
		if c == other {
			cs = append(cs[:i], cs[i+1:]...)
			return
		}
	}
}

type MinionCard struct {
	cost   int
	Attack int
	Health int
}

func NewMinionCard(attack, health int) *MinionCard {
	m := &MinionCard{Attack: attack, Health: health, cost: int(math.Ceil(float64(attack+health) / 2.0))}
	return m
}

func (m *MinionCard) Cost() int {
	return m.cost
}

func (mc *MinionCard) play(gs *gameState) {
	// deduct cost from available mana
	gs.pc.manaAvailable -= mc.cost

	// create a minion permanent on the current player's board
	mp := newMinionPerm(mc)
	gs.pc.permanents = append(gs.pc.permanents, mp)

	fmt.Fprintln(gs.logger, "Play card", mp)
}

func (m *MinionCard) String() string {
	return fmt.Sprintf("minionCard<%d:%d/%d>", m.cost, m.Attack, m.Health)
}
