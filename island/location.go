package island

import "fmt"

type Location struct {
	//TODO b Being
	n int
}

func (l *Location) String() string {
	return fmt.Sprintf("%d", l.n)
}
