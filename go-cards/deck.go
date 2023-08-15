package main

import (
	"fmt"
	"strings"
)

// type "deck" is slice of strings
type deck []string

// Reciever func
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardvalues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	for _, suit := range cardSuits {
		for _, cardValue := range cardvalues {
			cards = append(cards, cardValue+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //	:2 = 0:2 = begining to 1 (exclude 2) 2: = all after 2 include 2
}

func (d deck) deckToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) {
	// return io.WriteString(fileName, d.deckToString())
}
