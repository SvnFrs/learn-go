package random

import (
	"math/rand"
)

func CreateRandomNumbers() int {
	return rand.Intn(100)
}
