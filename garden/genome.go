package garden

import "math/rand"

const (
	maxByte           = 256
	cardsInDeck       = 30
	geneLength        = 3
	uniqueMarkerCount = 2
	headerLength      = 1
	genomeSize        = headerLength + cardsInDeck*geneLength
)

type Marker byte

// NOTE: be sure to update uniqueMarkerCount, above
const (
	MARKER_NULL Marker = iota
	MARKER_MINION_CARD
)

type Genome []byte
type GeneticCode []byte

func NewRandomGenome() Genome {
	bs := make([]byte, genomeSize)
	for i, _ := range bs {
		bs[i] = byte(rand.Intn(maxByte))
	}
	return bs
}

func (g Genome) Replicate() Genome {
	child := make([]byte, genomeSize)
	copy(child, g)

	mutations := rand.Intn(5)
	for im := 0; im < mutations; im++ {
		locus := rand.Intn(cardsInDeck * geneLength)
		child[locus] = byte(rand.Intn(maxByte))
	}

	return child
}

func (d *Deck) ToGenome() Genome {
	bs := make([]byte, 0)

	// Header
	bs = append(bs, d.Endurance)

	// Cards
	for _, c := range d.Cards {
		bs = append(bs, encodeCard(c)...)
	}
	return bs
}

func encodeCard(c Card) GeneticCode {
	switch v := c.(type) {
	case *MinionCard:
		return []byte{
			byte(MARKER_MINION_CARD),
			byte(v.Attack),
			byte(v.Health),
		}
	default:
		return []byte{
			byte(MARKER_NULL),
			0,
			0,
		}
	}
}

func (bs Genome) ToDeck() *Deck {
	d := new(Deck)
	d.Cards = make([]Card, 0)

	// Header
	d.Endurance = bs[0]

	// Cards
	for i := 1; i < len(bs); /*empty*/ {
		switch decodeMarker(bs[i]) {
		case MARKER_MINION_CARD:
			c, byteLen := decodeMinionCard(GeneticCode(bs[i:]))
			d.Cards = append(d.Cards, c)
			i += byteLen
		case MARKER_NULL:
			c := NewMinionCard(0, 0)
			d.Cards = append(d.Cards, c)
			i += geneLength
		default:
			i++
		}
	}

	return d
}

func decodeMinionCard(bs GeneticCode) (*MinionCard, int) {
	m := NewMinionCard(int(bs[1]%20), int(bs[2]%20+1))
	return m, geneLength
}

func decodeMarker(b byte) Marker {
	return Marker(b % uniqueMarkerCount)
}
