package main

import (
	"bufio"
	"literatorture/messages"
	"os"
)

var reader bufio.Reader

func init() {
	reader = *bufio.NewReader(os.Stdin)
}

func main() {

	running := true
	messages.PrintWelcomeMessage()

	for running {
		messages.PrintMenu()
		input, inputErr := readInt()
		if inputErr != nil {
			messages.PrintError(inputErr)
			continue
		}
		running = selectOption(input)
	}

	messages.PrintGoodbyeMessage()

}

// Selects the option which user has inputed.
// Returns a boolean which determines if the program will keep running or not
func selectOption(input int) bool {
	switch input {
	case 0:
		return false
	case 1:
		insertWord()
		return true
	default:
		messages.PrintInvalidOptionMessage()
		return true
	}
}

// Inserts a word and its meaning in the dictionary
func insertWord() {
	word, wordErr := readStringInput()
}
