package main

import "fmt"

func showGenres() {
	var genres [3]string = [...]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("genres[0]: %#v\n", genres[0])
	fmt.Printf("genres[1]: %#v\n", genres[1])
	fmt.Printf("genres[2]: %#v\n", genres[2])
	genres[1] = "Sci-Fi"
	fmt.Printf("genres[1]: %#v\n", genres[1])
	fmt.Printf("genres: %#v\n", genres)
}
 
func main() {
	showGenres()
}