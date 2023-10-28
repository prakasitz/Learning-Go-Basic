package main

import "fmt"

func main() {
	//Workshop:  Println

	var movieName string = "Avengers: Endgame"
	var year = 2019
	rating := 8.434
	var typeMovie string = `- Sci-Fi
	- Action`
	var isSuperHero bool = true

	fmt.Printf("เรื่อง: %s\n", movieName)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %.1f\n", rating)
	fmt.Printf("ประเภท: %s\n", typeMovie)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperHero)
	fmt.Printf("emoji: %c\n", '🎆')
}