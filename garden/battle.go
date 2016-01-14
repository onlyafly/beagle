package garden

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func takeTurnActions(gs *gameState) {
	for {
		actionCount := 0

		if gs.pc.manaAvailable > 0 {
			playableCards := gs.pc.hand.withinManaRange(0, gs.pc.manaAvailable)
			playableCards.Shuffle()

			if len(playableCards) > 0 {
				c := playableCards[0]
				gs.pc.hand.RemoveCard(c)

				c.play(gs)
				actionCount++
			}
		}

		actionablePerms := gs.pc.permanents.findActionablePerms()

		if len(actionablePerms) > 0 {
			perm := actionablePerms[0]
			perm.act(gs)
			actionCount++
		}

		// stop looping once there are no actions taken
		if actionCount == 0 {
			return
		}
	}
}

type gameState struct {
	turn   int
	logger io.Writer
	pc     *player // current player
	pe     *player // enemy player
}

func Battle(da *Deck, db *Deck) int {
	gs := new(gameState)
	gs.logger = os.Stdout
	ps := []*player{newPlayer("A", da), newPlayer("B", db)}

	// Prepare the players
	for _, p := range ps {
		p.library.Shuffle()
		p.drawCards(3)
	}

	firstPlayerNumber := rand.Intn(2)
	fmt.Fprintln(gs.logger, "Winner of flip:", firstPlayerNumber)

	for iTurn := firstPlayerNumber; iTurn < 100; iTurn++ {
		gs.turn = iTurn
		gs.pc = ps[iTurn%2]
		gs.pe = ps[(iTurn+1)%2]

		// Begin turn
		fmt.Fprintln(gs.logger, "Begin turn ", gs.pc)
		gs.pc.manaCrystals++
		gs.pc.manaAvailable = gs.pc.manaCrystals
		for _, perm := range gs.pc.permanents {
			perm.setActionable(true)
		}

		// Main phase
		takeTurnActions(gs)

		// End turn
		for i, p := range ps {
			if p.health <= 0 {
				// p has lost, return the other player's id
				return (i + 1) % 2
			}
		}
	}

	return -1
}
