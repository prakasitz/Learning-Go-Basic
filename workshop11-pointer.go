package main

import "fmt"

type movie1 struct {
	title	   string;
	year	   int;
	rating	   float32;
	votes	   []float64;
	genres	   []string;
	isSuperHero bool;
}

// method
func (mv *movie1) addVote(v float64) {
	// (*mv).votes = append((*mv).votes, v); // same as below 
	// because (*mv) automatically dereference the pointer in GO
	mv.votes = append(mv.votes, v);
}

func workshop() {
	// mv_addr := new(movie1); // same as &movie1{}

	var mv_addr *movie1 = &movie1{
		title: "Avengers: Endgame",
		year: 2019,
		rating: 8.4,
		votes: []float64{8.4, 8.5, 8.6},
		genres: []string{"Action", "Adventure", "Drama"},
		isSuperHero: true,
	};

	fmt.Printf("Before Add %v\n", mv_addr);
	mv_addr.addVote(10.0);
	fmt.Printf("After Add %v\n", mv_addr);
}

func basic_pointer() {
	var price float64 = 100.12;
	var addr *float64 = &price; // addr is a pointer to a float64
	fmt.Printf("1 value of price: %.2f\n", price);
	fmt.Printf("2 address of price in addr:  %v\n", addr); // address of price
	fmt.Printf("3 value of address in addr: %v\n", *addr); // value of address of addr
	fmt.Printf("5 address of addr: %v\n", &addr); // address of addr
	fmt.Printf("4 type of addr %T\n", addr); // type of addr
}

func main() {
	fmt.Println("================ BASIC ===================");
	basic_pointer();
	fmt.Println("================= WORKSHOP =================");
	workshop();
	
}