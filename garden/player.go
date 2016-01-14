package garden

type player struct {
	health        int
	manaCrystals  int
	manaAvailable int
	library       Cards
	permanents    []permanent
	hand          Cards
}

func newPlayer(d *Deck) *player {
	// Make l a copy of the cards in the deck
	l := make([]Card, len(d.Cards))
	copy(l, d.Cards)

	p := &player{
		health:  30,
		library: l,
	}

	return p
}

// Attempt to draw n cards. Returns the number of cards actually drawn
func (p *player) drawCards(n int) int {
	i := 0
	for ; i < n; i++ {
		p.hand = append(p.hand, p.library[0])
		p.library = p.library[1:]
	}
	return i
}
