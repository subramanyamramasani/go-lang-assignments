package main

import (
	"errors"
	"fmt"
)

func test(value int) (int, error) {

	if value == 0 {

		return 0, nil

	} else {

		return -1, errors.New("Invalid value: ")

	}
}

func main() {
	value, error := test(10)
	fmt.Printf("Value: %d, Error: %v", value, error)

	value, error = test(0)
	fmt.Printf("\n\nValue: %d, Error: %v", value, error)
}
