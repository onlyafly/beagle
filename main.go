package main

import (
	"fmt"
	"galapagos/island"
)

func main() {
	fmt.Println("Hello, playground")
	s := island.NewEcosystem()
	fmt.Println(s.Board)
}
