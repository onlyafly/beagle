package main

import (
	"bytes"
	"fmt"
	"galapagos/garden"
	"math/rand"
	"sort"
	"strconv"
)

const (
	ARENA_SIZE        = 1000
	ROUNDS            = 5000
	BATTLES_PER_ROUND = 3
)

var (
	originalArena []*competitor
	arena         []*competitor
	battleResults []*garden.BattleResult
)

type competitor struct {
	strainId   int
	generation int
	children   int
	lineage    []int
	g          garden.Genome
	d          *garden.Deck
	wins       int
	ties       int
	loses      int
}

func (c *competitor) String() string {
	mc := c.d.ManaCurve()
	var buffer bytes.Buffer
	for i := 0; i < 20; i++ {
		buffer.WriteString(strconv.Itoa(mc[i]) + " ")
	}

	return fmt.Sprintf("COMPETITOR:\n -- DECK: %v\n -- LINEAGE: %v\n -- GENERATION: %v\n -- W-T-L: %v-%v-%v (%v%%)\n -- CURVE: %v\n", c.d, c.lineage, c.generation, c.wins, c.ties, c.loses, float32(c.wins)/float32(c.loses+c.ties+c.wins)*100.0, buffer.String())
}

type ByRecord []*competitor

func (a ByRecord) Len() int      { return len(a) }
func (a ByRecord) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRecord) Less(i, j int) bool {
	return float32(a[i].wins)/float32(a[i].loses) < float32(a[j].wins)/float32(a[j].loses)
}

func seedArena() {
	fmt.Println("Seeding arena...")

	arena = make([]*competitor, ARENA_SIZE)
	for i := range arena {
		g := garden.NewRandomGenome()
		d := g.ToDeck()
		arena[i] = &competitor{
			lineage:  []int{i},
			strainId: i,
			g:        g,
			d:        d,
		}
	}

	originalArena = make([]*competitor, ARENA_SIZE)
	copy(originalArena, arena)
}

func shuffleArena() {
	// The Fisher-Yates shuffling algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	for i := range arena {
		j := rand.Intn(i + 1)
		arena[i], arena[j] = arena[j], arena[i]
	}
}

func reproduce(parent *competitor) *competitor {
	childGenome := parent.g.Replicate()
	childDeck := childGenome.ToDeck()
	childLineage := make([]int, len(parent.lineage)+1)
	copy(childLineage, parent.lineage)
	childLineage[len(childLineage)-1] = parent.children
	parent.children++
	childCompetitor := &competitor{
		lineage:    childLineage,
		generation: parent.generation + 1,
		strainId:   parent.strainId,
		g:          childGenome,
		d:          childDeck,
	}
	return childCompetitor
}

func pruneAndReproduce() {
	sort.Sort(ByRecord(arena)) // sort the arena, ascending by record

	// Remove the bottom 500 competitors
	arena = arena[len(arena)/2:]

	// Allow all remaining in arena to procreate
	arenaLen := len(arena)
	for i := 0; i < arenaLen; i++ {
		c := arena[i]

		// If a competitor has played more games than its endurance, it dies and an offspring takes its place
		if c.wins > int(c.d.Endurance) {
			arena[i] = reproduce(c)
		}

		arena = append(arena, reproduce(c))
	}
}

func runRound(round int) {
	// Prepare for this round run
	battleResults = make([]*garden.BattleResult, 0)

	// Execute this round's battles
	for ibattle := 0; ibattle < BATTLES_PER_ROUND; ibattle++ {
		runBattle()
	}
}

func runBattle() {
	shuffleArena()

	splitPoint := len(arena) / 2
	for i := 0; i < splitPoint; i++ {
		c0 := arena[i]
		c1 := arena[i+splitPoint]

		battleResult := garden.Battle(c0.d, c1.d, &junk{})
		battleResults = append(battleResults, battleResult)

		switch battleResult.Winner {
		case 0:
			c0.wins++
			c1.loses++
		case 1:
			c0.loses++
			c1.wins++
		case -1:
			c0.ties++
			c1.ties++
		}
	}
}

type statResults struct {
	competitorWithMostWins *competitor
	avgTurns               float32
	countCompetitors       int
}

func (sr *statResults) String() string {
	return fmt.Sprintf("STATS:\n -- #COMPETITORS %v\n -- AVG TURNS: %v\n -- TOP COMPETITOR: %v", sr.countCompetitors, sr.avgTurns, sr.competitorWithMostWins)
}

func stats() *statResults {
	iCompetitorWithMostWins := 0
	mostWins := 0

	for i, c := range arena {
		if c.wins > mostWins {
			iCompetitorWithMostWins = i
			mostWins = c.wins
		}
	}

	totalTurns := 0
	for _, r := range battleResults {
		totalTurns += r.Turns
	}
	averageTurns := float32(totalTurns) / float32(len(battleResults))

	return &statResults{
		competitorWithMostWins: arena[iCompetitorWithMostWins],
		avgTurns:               averageTurns,
		countCompetitors:       len(arena),
	}
}
