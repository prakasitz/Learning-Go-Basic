package main

import "fmt"

type movie struct {
	title       string;
	year        int;
	rating      float64;
	genres      []string;
	isSuperHero bool;
}

func (mv movie) info() {
	fmt.Printf("%s (%d) - %.2f\n", mv.title, mv.year, mv.rating);
	
	fmt.Println("Genres:")
	for _, genre := range mv.genres {
		fmt.Printf("\t\t%s\n", genre);
	}
}

func main() {

	ae := movie{
		title: "Avengers: Endgame",
		year: 2019,
		rating: 8.4,
		genres: []string{"Action", "Adventure", "Drama"},
		isSuperHero: true,
	};

	ae.info();

}