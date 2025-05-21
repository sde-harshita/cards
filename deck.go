package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// converts a deck >> []string >> string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// writes the deck cards to a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// reads the deck cards from a file
func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// shuffles the deck of cards using a random number generation
func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		p := r.Intn(len(d) - 1)
		d[i], d[p] = d[p], d[i]
	}
}
