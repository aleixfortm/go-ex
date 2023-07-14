package main

import "fmt"

// this new type implements a function that is applicable to any type of the program that have de getGreeting function and return a string from it
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// Contained logic is completely different from the same function for the spanishBot (fully independent)
	return "Hey there!"
}

func (sb spanishBot) getGreeting() string {
	// Contained logic is completely different from the same function for the englishBot (fully independent)
	return "Hola, que tal?"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
