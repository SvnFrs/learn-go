package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)

	fmt.Println("How many people are there?")
	var number int
	fmt.Scanf("%d", &number)

	var name []string
	for i := 0; i < number; i++ {
		fmt.Println("Please enter your name!")
		var input string
		fmt.Scanf("%s", &input)
		name = append(name, input)
	}
	message, err := greetings.Hellos(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
