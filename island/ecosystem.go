package island

type Ecosystem struct {
	Board Board
}

func NewEcosystem() *Ecosystem {
	i := new(Ecosystem)
	i.Board = NewBoard()

	return i
}
