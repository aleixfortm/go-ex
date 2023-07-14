package main

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

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
