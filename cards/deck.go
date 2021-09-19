package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck
// which is a slice of strings

type deck []string

// creates and return a list of cards
// no receiver needed because we are creating a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

// returns the hand and the remaining part of the deck
// Multiple return values here.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// takes a deck and makes it a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
