package main

import "fmt"

func main() {
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "Queen of Hearts")
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.deckToString())
}

// func newCard() string {
// 	return "Ace of Diamonds"
// }
