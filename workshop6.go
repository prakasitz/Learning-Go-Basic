package main

import "fmt"

func showGenres() {
	var genres [3]string = [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genres)
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie: " + genres[i]
	}
	fmt.Printf("after for loop: %#v\n", genres)
	
	for i, genre := range genres {
		fmt.Printf("genres[%d]: %#v\n", i, genre)
	}
}
 
func main() {
	showGenres()
}