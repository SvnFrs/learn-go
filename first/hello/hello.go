package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)

	fmt.Println("Please enter your name!")
	var name string
	fmt.Scanf("%v", &name)

	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
