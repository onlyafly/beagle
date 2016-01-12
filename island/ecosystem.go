package island

import "fmt"

type Ecosystem struct {
	Board   Board
	Turtles []*Turtle
}

func NewEcosystem() *Ecosystem {
	ec := new(Ecosystem)
	ec.Board = NewBoard()
	return ec
}

func (ec *Ecosystem) AddTurtle(x int, y int) *Turtle {
	t := NewTurtle()

	ec.Turtles = append(ec.Turtles, t)

	l := ec.Board[x][y]
	l.Turtle = t
	t.X = x
	t.Y = y

	return t
}

func (ec *Ecosystem) Tick() {
	for _, t := range ec.Turtles {
		t.Age++

		inst := t.NextInstruction()

		switch inst {
		case NOTHING:
		case REPRODUCE:
			if t.Energy > 100 {
				cx, cy := ec.Board.EmptyNeighbor(t.X, t.Y)
				fmt.Println("tx,ty,cx,cy = ", t.X, t.Y, cx, cy)

				if cx == -1 {
					// there are no empty neighbors, so reduce the energy due to overcrowding
					t.Energy = t.Energy - 10
				} else {
					// Create child
					ct := ec.AddTurtle(cx, cy)
					ct.Energy = t.Energy / 2
					ct.Data = t.Data

					// Update ParseEvalPrint
					t.Energy = t.Energy / 2
				}
			}
		case LEECH:
			t.Energy++
		}
	}
}
