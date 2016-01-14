package main

import (
	"fmt"
	"galapagos/garden"
)

func main() {
	da := garden.RandomDeck()
	db := garden.RandomDeck()

	fmt.Println(da)
	fmt.Println(db)

	r := garden.Battle(da, db)

	fmt.Println("AND THE RESULT IS...")
	fmt.Println(r)
}
