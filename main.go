package main

import "fmt"

func main() {
	// Array has fixed size. In this example it's 3.
	// cards := [3]string{"Ace of Diamonds", newCard()}
	// cards[2] = "Six of Spades"
	// for i := 0; i < len(cards); i++ {
	// 	fmt.Println(i, cards[i])
	// }

	// Slice could grow or shrink
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
