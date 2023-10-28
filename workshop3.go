package main

import (
	"fmt"
)

func main() {
	var rating float64 = 6.5
	fmt.Printf("rating: %.1f\n", rating)

	if rating < 5.0 {
		fmt.Println("Disappointing 😞")
	} else if rating < 7.0 {
		fmt.Println("Normal 😐")
	} else if rating < 9.0 {
		fmt.Println("Good 😊")
	} else {
		fmt.Println("👽👻👻👻👻")
	}
}