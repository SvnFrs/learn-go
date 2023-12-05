package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("consider writing a proper name next time")
	}

	message := fmt.Sprintf("Hi %v, Welcome to Go!", name)
	return message, nil
}
