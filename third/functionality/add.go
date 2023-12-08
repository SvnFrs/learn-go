package functionality

import "fmt"

var a int
var b int

func Add() int {
	fmt.Println("Please enter A and B")

	fmt.Printf("A: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("\nB: ")
	fmt.Scanf("%d", &b)

	result := do(a, b)

	return result
}

func do(a int, b int) int {
	return a + b
}
