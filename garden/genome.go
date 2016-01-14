package garden

type Marker byte

const (
	geneLength        = 3
	uniqueMarkerCount = 2
)

// NOTE: be sure to update uniqueMarkerCount, above
const (
	MARKER_NULL Marker = iota
	MARKER_MINION_CARD
)

func EncodeDeck(d *Deck) []byte {
	bs := make([]byte, 0)
	for _, c := range d.Cards {
		bs = append(bs, EncodeCard(c)...)
	}
	return bs
}

func EncodeCard(c Card) []byte {
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

func DecodeDeck(bs []byte) *Deck {
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

func DecodeMinionCard(bs []byte) (*MinionCard, int) {
	m := NewMinionCard(int(bs[1]), int(bs[2]))
	return m, geneLength
}

func decodeMarker(b byte) Marker {
	return Marker(b % uniqueMarkerCount)
}
