package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/divan/num2words"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Jack", "Queen", "Knight"}

	for i := 2; i < 11; i++ {
		cardValues = append(cardValues, strings.Title(num2words.Convert(i)))
	}

	for i, value := range cardValues {
		for j, suit := range cardSuits {
			cards = append(cards, fmt.Sprintf("[%d] ", len(cardSuits)*i+(j+1))+value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), os.FileMode(0666))
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Error handling:
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), "\n"))
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
