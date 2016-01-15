package main

import (
	"fmt"
	"galapagos/garden"
)

func main() {
	ga := garden.RandomGenome()
	gb := garden.RandomGenome()
	da := garden.DecodeDeck(ga)
	db := garden.DecodeDeck(gb)

	fmt.Println(da)
	fmt.Println(db)

	r := garden.Battle(da, db)

	fmt.Println("AND THE RESULT IS...")
	fmt.Println(r)
}
