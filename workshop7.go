package main

import "fmt"

func showVote(v []float64) {
	fmt.Printf("vote: %#v\n", v)
}

func main() {
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	var votes []float64

	for _, x := range xs {
		votes = append(votes, x)
	}

	for _, y := range ys {
		votes = append(votes, y)
	}
	aa := append(xs, ys...)
	showVote(aa)

	showVote(votes)

	if(len(votes) > 9) {
		fmt.Println("votes 5 - 8:", votes[5:9])
	}
}