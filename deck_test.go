package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deckLength := 52
	d := newDeck()
	if len(d) != deckLength {
		t.Errorf("Expected deck length of %d, but got %d", deckLength, len(d))
	}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	eachSuitCardsAmount := deckLength / len(cardSuits)
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
	if d[len(d)-1] != fmt.Sprintf("[%d] Ten of Clubs", deckLength) {
		t.Errorf("Last card in deck should be \"[%d] Ten of Clubs\", but got \"%s\"", deckLength, d[len(d)-1])
	}

}
