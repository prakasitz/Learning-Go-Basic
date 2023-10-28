package main

import "fmt"

func show(tag string, s []int) {
	l := len(s)
	c := cap(s)
	fmt.Printf("%v: len=%d cap=%d %v\n", tag, l, c, s)
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -2}
	show("slice1", slice1)

	s1 := slice1[0:2]
	show("s1", s1)

	s2 := slice1[1:12]
	show("s2", s2)

	_ = append(s2, 99)
	show("s2", s2)

	show("slice1", slice1)

}

