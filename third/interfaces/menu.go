package interfaces

import (
	"fmt"
)

func Menu() {
	fmt.Println("Welcome to the Go calculator")
	fmt.Println("1. Add")
	fmt.Println("2. Substract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Print("\tPlease enter your choice: ")
}
