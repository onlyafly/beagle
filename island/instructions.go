package island

type Instruction byte

const (
	NOTHING Instruction = iota
	REPRODUCE
	LEECH
)
