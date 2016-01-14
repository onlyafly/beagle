package main

import (
	"fmt"
	"galapagos/garden"
)

func main() {
	da := garden.RandomDeck()
	db := garden.RandomDeck()

	da.Cards.Shuffle()

	fmt.Println(da)
	fmt.Println(db)

	//TODO r := game.Battle(da, db)

	fmt.Println("AND THE RESULT IS...")
	//TODO fmt.Println(r)
}
