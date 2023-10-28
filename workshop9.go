package main

import "fmt"

func main() {

	type movie struct {
		title       string;
		year        int;
		rating      float64;
		genres      []string;
		isSuperHero bool;
	}

	var movie1 = movie{
		title: "Avengers: Endgame", 
		year: 2019, 
		rating: 8.4, 
		genres: []string{"Action", "Adventure", "Drama"}, 
		isSuperHero: true,
	};
	
	var movie2 = movie{
		title: "Avengers: Infinity War", 
		year: 2018, 
		rating: 8.5, 
		genres: []string{"Action", "Adventure", "Drama"}, 
		isSuperHero: true,
	};

	var mvs []movie;
	mvs = append(mvs, movie1, movie2);
	
	for _, mv := range mvs {
		fmt.Printf("main.movie{title: %s, year: %d, rating: %.2f, genres: %v, isSuperHero: %t}\n", mv.title, mv.year, mv.rating, mv.genres, mv.isSuperHero);
	}
}