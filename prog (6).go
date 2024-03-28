package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// print deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// slice of string
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// string t0 vectors
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save file form hard drive
func (d deck) saveFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read from  hard drive
// Error handling
func readFromHD(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

//Random generation
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		temp := r.Intn(len(d) - 1)
		d[i], d[temp] = d[temp], d[i]
	}
}

func main() {
	// cards := newdeck()
	// hands, remainDeck := deal(cards, 5)
	// hands.print()
	// remainDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveFile("my_cards")
	// cards:=readFromHD("my_cards")
	// cards.print()
	cards := newdeck()
	cards.shuffle()
	cards.print()
}
