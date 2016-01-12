package island

type Ecosystem struct {
	Board   Board
	Turtles []*Turtle
}

func NewEcosystem() *Ecosystem {
	i := new(Ecosystem)
	i.Board = NewBoard()

	return i
}

func (ec *Ecosystem) AddTurtle(x int, y int) {
	t := new(Turtle)
	t.n = 42

	ec.Turtles = append(ec.Turtles, t)

	l := ec.Board[x][y]
	l.Turtle = t
}

func (ec *Ecosystem) Tick() {
	for _, t := range ec.Turtles {
		t.Age++
		t.Energy--
	}
}
