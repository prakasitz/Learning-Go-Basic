package main

import "fmt"

func main() {
	//Workshop:  Println

	var movieName string = "Avengers: Endgame"
	var year = 2019
	rating := 8.4
	var typeMovie string = `- Sci-Fi
	 - Action`
	var isSuperHero bool = true

	fmt.Println("เรื่อง: ", movieName)
	fmt.Println("ปี: ", year)
	fmt.Println("เรตติ้ง: ", rating)
	fmt.Println("ประเภท: ", typeMovie)
	fmt.Println("ซุปเปอร์ฮีโร่: ", isSuperHero)
}