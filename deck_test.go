package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, rc := deal(5, d)

	if len(hand) != 5 {
		t.Errorf("Expected a hand of 5 cards, but got %v", len(hand))
	}

	if len(rc) != 47 {
		t.Errorf("Expected 47 cards in remaining deck, but got %v", len(rc))
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	_d := readFromFile("_decktesting")

	if len(_d) != 52 {
		t.Errorf("Expected 52 cards in the deck, but got %v", len(_d))
	}

	os.Remove("_decktesting")
}
