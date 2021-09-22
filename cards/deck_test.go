package main

import "testing"

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
