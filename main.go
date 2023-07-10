package main

import "fmt"

func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	cards, err := readFromFile("my_cards")
	if err == nil {
		fmt.Println(cards)
	}
}
