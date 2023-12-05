package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println("Please enter your name!")
	var name string
	fmt.Scanf("%v", &name)

	message := greetings.Hello(name)
	fmt.Println(message)
}
