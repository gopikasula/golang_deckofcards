package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected num of cards to be 16, but found %v", len(d))
	}
}

func TestSaveDeckToFileMakeDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	readDeck := newDeckFromFile("_decktesting")

	if len(readDeck) != len(d) {
		t.Errorf("Expected 16 cards in deck, got %v", len(readDeck))
	}

	os.Remove("_decktesting")
}
