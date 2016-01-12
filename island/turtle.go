package island

import "strconv"

type Turtle struct {
	Energy             int
	Age                int
	Data               []Instruction
	CurrentInstruction int
	X                  int
	Y                  int
}

func NewTurtle() *Turtle {
	t := new(Turtle)
	t.Data = []Instruction{REPRODUCE, LEECH}
	return t
}

func (t *Turtle) String() string {
	return strconv.Itoa(t.Energy)
}

func (t *Turtle) NextInstruction() Instruction {
	if t.CurrentInstruction >= len(t.Data) {
		t.CurrentInstruction = 0
	}
	i := t.Data[t.CurrentInstruction]
	t.CurrentInstruction++
	return i
}
