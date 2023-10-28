package main

import "fmt"

func emote(rating float64) string {
	// convert to switch case
	switch {
	case rating < 5.0:
		return "Disappointing 😞"
	case rating < 7.0:
		return "Normal 😐"
	case rating < 9.0:
		return "Good 😊"
	default:
		return "👽👻👻👻👻"
	}
}
 
func main() {
	fmt.Println(emote(6.5))
	fmt.Println(emote(9.5))
	fmt.Println(emote(3.5))
	fmt.Println(emote(7.5))
}