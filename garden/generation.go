package garden

func RandomDeck() *Deck {
	d := new(Deck)
	d.Cards = make([]Card, 30)
	for i, _ := range d.Cards {
		m := NewMinionCard(i+1, i+1)
		d.Cards[i] = m
	}
	return d
}
