package main

import "fmt"


func main() {
	var msg string = "hello world"
	var age int = 30
	price, isFound := 100.25, true
	var r rune = '🎆'

	fmt.Println("msg: ", msg)
	fmt.Println("age: ", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("isFound: %t\n", isFound)
	fmt.Printf("r: %c\n",r)

	fmt.Printf("price: %#v\n", price)
	fmt.Printf("isFound: %#v\n", isFound)
	fmt.Printf("r: %c\n",r)

	fmt.Printf("type: %T -- price: %#v\n", price, price)
	fmt.Printf("type: %T -- isFound: %#v\n", isFound, isFound)
}