package garden

type Deck struct {
	Cards Cards
}

func (d *Deck) ManaCurve() map[int]int {
	m := make(map[int]int)
	for _, c := range d.Cards {
		m[c.Cost()]++
	}
	return m
}
