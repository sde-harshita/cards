package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Five of Hearts")
	cards.print()
}

func newCard() string {
	return "King of Spades"
}
