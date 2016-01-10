package main

import (
	"beagle/island"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	myIsland := island.NewIsland()
	fmt.Println(myIsland.Board)
}
