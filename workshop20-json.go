package main

import (
	"encoding/json"
	"fmt"
)



type user struct {
	Name string `json:"name"`;
	Age int	`json:"age"`;
	Hobbies []string `json:"hobbies"`;
}

func main() {
	//user json string
	user_json := `[{"name":"John Doe","age":25,"hobbies":["Reading","Gaming","Coding"]}]`;

	//decode user_json with json package 
	userA := []user{};
	err := json.Unmarshal([]byte(user_json), &userA);

	if err != nil {
		fmt.Println(err);
	}
	//
	fmt.Printf("User: %#v\n", userA)
	
	

}