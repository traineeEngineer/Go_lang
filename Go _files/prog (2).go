package main

import (
	"fmt"
	"strings"
)

type deck []string

func newdeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardsSuits {
		for _, value := range cardsValue {
			cards = append(cards, suits+"of"+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func main() {
	cards := newdeck()
	// hands, remainDeck := deal(cards, 5)
	// hands.print()
	// remainDeck.print()

	fmt.Println(cards.toString())
}
