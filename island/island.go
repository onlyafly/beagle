package island

type Island struct {
	Board Board
}

func NewIsland() *Island {
	i := new(Island)
	i.Board = NewBoard()

	return i
}
