package island

import "fmt"

type Location struct {
	//TODO b Being
	n      int
	Turtle *Turtle
}

func (l *Location) String() string {
	return fmt.Sprintf("%d", l.n)
}
