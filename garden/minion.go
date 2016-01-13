package garden

import (
	"fmt"
	"math"
)

type Card interface {
	Cost() int
}

type Minion struct {
	cost   int
	Attack int
	Health int
}

func NewMinion(attack, health int) *Minion {
	m := &Minion{Attack: attack, Health: health, cost: int(math.Ceil(float64(attack+health) / 2.0))}
	return m
}

func (m *Minion) Cost() int {
	return m.cost
}

func (m *Minion) String() string {
	return fmt.Sprintf("<M%d:%d/%d>", m.cost, m.Attack, m.Health)
}
