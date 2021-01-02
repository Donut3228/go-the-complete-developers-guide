package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type Config struct {
	DeckLength int
}

func newConfig() Config {
	return Config{DeckLength: 52}
}

func TestNewDeck(t *testing.T) {
	c := newConfig()
	d := newDeck()
	if len(d) != c.DeckLength {
		t.Errorf("Expected deck length of %d, but got %d", c.DeckLength, len(d))
	}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	eachSuitCardsAmount := c.DeckLength / len(cardSuits)
	for _, suit := range cardSuits {
		count := 0
		for _, card := range d {
			if strings.HasSuffix(card, suit) {
				count++
			}
		}
		if eachSuitCardsAmount != count {
			t.Errorf("Expected number of %s is %d, but got %d", suit, eachSuitCardsAmount, count)
		}
	}
	if d[0] != "[1] Ace of Spades" {
		t.Errorf("First card in deck should be \"Ace of Spades\", but got \"%s\"", d[0])
	}
	if d[len(d)-1] != fmt.Sprintf("[%d] Ten of Clubs", c.DeckLength) {
		t.Errorf("Last card in deck should be \"[%d] Ten of Clubs\", but got \"%s\"", c.DeckLength, d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	c := newConfig()
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != c.DeckLength {
		t.Errorf("Expected %d cards in deck, got %d", c.DeckLength, len(loadedDeck))
	}
	os.Remove("_decktesting")
}
