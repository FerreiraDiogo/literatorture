package input

import (
	"bufio"
	"fmt"
	"literatorture/validation"
)

// Reads user input. This functions expects an int to be typed.
// Returns the inputed int and an error if the inputed value is not assignable
func ReadInt() (int, error) {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return 0, err
	}
	return input, nil
}

// Reads user input. This functions expects that a string is typed.
// Returns the inputed string and an error if the inputed value is not assignable
func ReadStringInput(reader *bufio.Reader) (string, error) {
	input, inputErr := reader.ReadString('\n')
	if inputErr != nil {
		return "", inputErr
	}
	validationErr := validation.IsValidWord(input)
	if validationErr != nil {
		return "", validationErr
	}
	return input, nil
}
