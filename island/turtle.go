package island

import "strconv"

type Turtle struct {
	n      int // TODO remove
	Energy int
	Age    int
}

func (t *Turtle) String() string {
	return strconv.Itoa(t.Age)
}
