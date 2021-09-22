package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck length of 40 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck first item to be Ace of Spades instead received %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected deck first item to be King of Clubs instead received %v", d[len(d)-1])
	}
}

// when dealing with file creation we need to make sure we get rid of the file we create
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected deck length of 40 but got %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected deck first item to be Ace of Spades instead received %v", d[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Expected deck first item to be King of Clubs instead received %v", d[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")

}
