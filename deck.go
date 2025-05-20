package main

import "fmt"

type deck []string

// prints cards present in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

// creats a new deck containing all cards and return the newly created deck
func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// draws a hand from the deck and returns the drawed hand and the remaining cards of the deck
func deal(handSize int, d deck) (deck, deck) {
	return d[:handSize], d[handSize:]
}
