package runner

import (
	"fmt"

	"example.com/functionality"
	interfaces "example.com/interface"
)

func Run() {
	validChoices := 5
	var choice int
	for {
		interfaces.Menu()
		fmt.Scanf("%d", &choice)
		if choice > validChoices {
			interfaces.Menu()
		} else {
			break
		}
	}
	switch choice {
	case 1:
		result := functionality.Add()
		fmt.Println("The result is: ", result)
	}

}
