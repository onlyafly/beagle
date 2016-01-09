package main

import (
	"beagle/island"
	"fmt"
)

var board island.Board

func main() {
	fmt.Println("Hello, playground")
	board := island.NewBoard()

	fmt.Println(board)
}
