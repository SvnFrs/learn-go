package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("consider writing a proper name next time")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hi, %v, Welcome to Go!",
		"Good luck on your journey, %v!",
		"Hail %v, well done.",
	}
	return formats[rand.Intn(len(formats))]
}
