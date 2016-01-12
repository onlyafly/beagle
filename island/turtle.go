package island

import "strconv"

type Turtle struct {
	n int
}

func (t *Turtle) String() string {
	return strconv.Itoa(t.n)
}
