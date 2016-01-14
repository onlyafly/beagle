package garden

import "math/rand"

func takeTurnActions(gs *gameState) {
	for {
		if gs.pc.manaAvailable > 0 {
			playableCards := cardsWithinManaRange(gs.pc.hand, 0, gs.pc.manaAvailable)
			playableCards.Shuffle()

			c := playableCards[0]
			gs.pc.hand.RemoveCard(c)

			c.play(gs)
		}
	}
}

type gameState struct {
	turn int
	pc   *player // current player
	pe   *player // enemy player
}

func Battle(da *Deck, db *Deck) int {
	gs := new(gameState)
	ps := []*player{newPlayer(da), newPlayer(db)}

	// Prepare the players
	for _, p := range ps {
		p.library.Shuffle()
		p.drawCards(3)
	}

	firstPlayerNumber := rand.Intn(2)

	for iTurn := firstPlayerNumber; ; iTurn++ {
		gs.turn = iTurn
		gs.pc = ps[iTurn%2]
		gs.pe = ps[(iTurn+1)%2]

		gs.pc.manaCrystals++
		gs.pc.manaAvailable = gs.pc.manaCrystals

		takeTurnActions(gs)
	}

	return -1
}
