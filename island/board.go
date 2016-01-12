package island

import "strings"

type Board [][]*Location

const boardSize = 10

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

func (b Board) Empty(x, y int) bool {
	return b[x][y].Turtle == nil
}

func (b Board) EmptyNeighbor(x, y int) (int, int) {
	left := x - 1
	if left < 0 {
		left = boardSize - 1
	}
	right := x + 1
	if right >= boardSize {
		right = 0
	}
	above := y - 1
	if above < 0 {
		above = boardSize - 1
	}
	below := x + 1
	if below >= boardSize {
		below = 0
	}

	//fmt.Println("left, x, right = ", left, x, right)
	//fmt.Println("above, y, below = ", above, y, below)

	for _, i := range []int{left, x, right} {
		for _, j := range []int{above, y, below} {
			if b.Empty(i, j) {
				return i, j
			}
		}
	}

	return -1, -1
}

func NewBoard() Board {
	ySize := boardSize
	xSize := boardSize

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
