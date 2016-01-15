package main

import (
	"fmt"
	"galapagos/garden"
	"math/rand"
)

type junk struct{}

func (j *junk) Write(p []byte) (n int, err error) {
	return 0, nil
}

/*
func main() {
	ga := garden.RandomGenome()
	gb := garden.RandomGenome()
	da := garden.DecodeDeck(ga)
	db := garden.DecodeDeck(gb)

	fmt.Println(da)
	fmt.Println(db)

	r := garden.Battle(da, db, &junk{})

	fmt.Println("AND THE RESULT IS...")
	fmt.Println(r)
}
*/

type competitor struct {
	g    garden.Genome
	wins int
}

func (c *competitor) String() string {
	return fmt.Sprintf("(%v won %v)", garden.DecodeDeck(c.g), c.wins)
}

type statResults struct {
	competitorWithMostWins *competitor
}

func (sr *statResults) String() string {
	return fmt.Sprintf("(RESULTS: %v)", sr.competitorWithMostWins)
}

func stats(arena []*competitor) *statResults {
	iCompetitorWithMostWins := 0
	mostWins := 0

	for i, c := range arena {
		if c.wins > mostWins {
			iCompetitorWithMostWins = i
			mostWins = c.wins
		}
	}

	return &statResults{
		competitorWithMostWins: arena[iCompetitorWithMostWins],
	}
}

func main() {
	arena := make([]*competitor, 1000)
	for i := range arena {
		arena[i] = &competitor{
			g: garden.RandomGenome(),
		}
	}

	for generation := 0; generation < 1000; generation++ {

		splitPoint := len(arena) / 2
		for ic1 := 0; ic1 < splitPoint; ic1++ {
			c1 := arena[ic1]
			c2 := arena[ic1+splitPoint]
			da := garden.DecodeDeck(c1.g)
			db := garden.DecodeDeck(c2.g)

			battleResult := garden.Battle(da, db, &junk{})

			winnerIndex := 0
			loserIndex := 0
			if battleResult.Winner == 0 {
				winnerIndex = ic1
				loserIndex = ic1 + splitPoint
			} else {
				winnerIndex = ic1 + splitPoint
				loserIndex = ic1
			}

			winner := arena[winnerIndex]
			winner.wins++
			arena[loserIndex] = &competitor{
				g: garden.Reproduce(winner.g),
			}
		}

		// The Fisher-Yates shuffling algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
		for i := range arena {
			j := rand.Intn(i + 1)
			arena[i], arena[j] = arena[j], arena[i]
		}

		// Show stats for this generation
		if generation%10 == 0 {
			fmt.Println("GENERATION", generation)
			fmt.Println(stats(arena))
		}
	}

	// Show the one with the most wins

	fmt.Println("AND THE RESULT IS...")
	fmt.Println(stats(arena))
}
