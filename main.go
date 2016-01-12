package main

import (
	"fmt"
	"galapagos/island"
)

func main() {
	fmt.Println("Hello, playground")

	s := island.NewEcosystem()

	s.AddTurtle(1, 1)

	fmt.Println(s.Board)
}
