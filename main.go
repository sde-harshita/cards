package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(5, cards)
	fmt.Println("\n Drawn Hand : ")
	hand.print()
	fmt.Println("\n Remaining Cards : ")
	remainingCards.print()
}
