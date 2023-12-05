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
	// random the text and print of the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// create a messages map
	messages := make(map[string]string)

	// use blank identifier to avoid redundancy of index value
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v, Welcome to Go!",
		"Good luck on your journey, %v!",
		"Hail %v, well done.",
	}
	return formats[rand.Intn(len(formats))]
}
