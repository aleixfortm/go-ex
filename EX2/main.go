package main

import "fmt"

func main() {
	slice := []int{}
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
	for i := range slice {
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
}
