package garden

import "fmt"

type player struct {
	name          string
	health        int
	manaCrystals  int
	manaAvailable int
	library       Cards
	permanents    permanents
	hand          Cards
}

func newPlayer(name string, d *Deck) *player {
	// Make l a copy of the cards in the deck
	l := make([]Card, len(d.Cards))
	copy(l, d.Cards)

	p := &player{
		name:    name,
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

func (p *player) String() string {
	return fmt.Sprintf("player<'%v',health=%d,crystals=%d>", p.name, p.health, p.manaCrystals)
}
