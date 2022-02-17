package main

import (
	"fmt"
	"log"

	"github.com/enikolas/cartographers/deck"
)

func main() {
	d, err := deck.NewDeck()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deck size: %d\n", len(d))

	for _, card := range d {
		fmt.Printf("\n---------\n\n%s\n", card.String())
	}
}
