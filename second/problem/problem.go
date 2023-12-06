package main

import (
	"fmt"
	"math"
)

func main() {
	var problems float32
	fmt.Println("How many problem do you have, I'll cut it in half")
	fmt.Scanf("%g", &problems)
	fmt.Printf("Now you have %g problems. Thank me later\n", math.Sqrt(float64(problems)))
}
