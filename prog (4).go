package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

// create new deck of cards
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

//print deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
//slice of string 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
// string t0 vectors
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
//save file form hard drive
func (d deck) saveFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func main() {
	cards := newdeck()
	// hands, remainDeck := deal(cards, 5)
	// hands.print()
	// remainDeck.print()

	// fmt.Println(cards.toString())
	cards.saveFile("my_cards")
}
