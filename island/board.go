package island

import "strings"

type Board [][]*Location

func (b Board) String() string {
	s := ""
	for i := range b {
		for j := range b[i] {
			l := b[i][j]
			if l.Turtle != nil {
				s += l.Turtle.String() + " "
			} else {
				s += "- "
			}
		}
		s = strings.TrimSpace(s)
		s += "\n"
	}

	return s
}

func NewBoard() Board {
	ySize := 10
	xSize := 10

	// Allocate the top-level slice
	board := make([][]*Location, ySize)

	// Allocate one large slice to hold all the locations
	locations := make([]*Location, xSize*ySize)
	for i := range locations {
		locations[i] = &Location{n: 0}
	}

	// Loop over the rows, slicing each row from the front of the remaining locations slice
	for i := range board {
		board[i], locations = locations[:xSize], locations[xSize:]
	}

	return board
}
