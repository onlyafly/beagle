package garden

type Deck struct {
	Cards     Cards
	Endurance byte
}

/*
func NewRandomDeck() *Deck {
	d := new(Deck)
	d.Cards = make([]Card, 30)
	for i, _ := range d.Cards {
		m := NewMinionCard(i+1, i+1)
		d.Cards[i] = m
	}
	return d
}
*/

func (d *Deck) ManaCurve() map[int]int {
	m := make(map[int]int)
	for _, c := range d.Cards {
		m[c.Cost()]++
	}
	return m
}
