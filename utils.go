package main

import (
	"fmt"
)

func chooseInt(minI int, maxI int) (int, error) {
	correct := false
	var input int
	for !correct {
		fmt.Print("> ")
		fmt.Scan(&input)
		if minI <= input && input <= maxI {
			correct = true
			break
		}
	}
	fmt.Println()
	return input, nil
}

