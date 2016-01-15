package garden

import "math/rand"

const (
	maxByte           = 256
	cardsInDeck       = 30
	geneLength        = 3
	uniqueMarkerCount = 2
)

type Marker byte
type Genome []byte

// NOTE: be sure to update uniqueMarkerCount, above
const (
	MARKER_NULL Marker = iota
	MARKER_MINION_CARD
)

func Reproduce(g Genome) Genome {
	child := make([]byte, cardsInDeck*geneLength)
	copy(child, g)

	mutations := rand.Intn(5)
	for im := 0; im < mutations; im++ {
		locus := rand.Intn(cardsInDeck * geneLength)
		child[locus] = byte(rand.Intn(maxByte))
	}

	return child
}

func RandomGenome() Genome {
	bs := make([]byte, cardsInDeck*geneLength)
	for i, _ := range bs {
		bs[i] = byte(rand.Intn(maxByte))
	}
	return bs
}

func EncodeDeck(d *Deck) Genome {
	bs := make([]byte, 0)
	for _, c := range d.Cards {
		bs = append(bs, EncodeCard(c)...)
	}
	return bs
}

func EncodeCard(c Card) Genome {
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

func DecodeDeck(bs Genome) *Deck {
	d := new(Deck)
	d.Cards = make([]Card, 0)

	for i := 0; i < len(bs); /*empty*/ {
		switch decodeMarker(bs[i]) {
		case MARKER_MINION_CARD:
			c, byteLen := DecodeMinionCard(bs[i:])
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

func DecodeMinionCard(bs Genome) (*MinionCard, int) {
	m := NewMinionCard(int(bs[1]), int(bs[2]))
	return m, geneLength
}

func decodeMarker(b byte) Marker {
	return Marker(b % uniqueMarkerCount)
}
