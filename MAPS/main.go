package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FFFF",
	}

	colors2 := make(map[int]string)
	colors2[1] = "#FFFFFF"
	colors2[2] = "#000000" // Accessing a value in maps works differently than structs
	// here we use square brackets instead

	// deleting key-value pairs from a map
	delete(colors, "red")
	delete(colors2, 1)

	fmt.Println(colors)
	fmt.Println(colors2)
}
