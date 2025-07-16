package main

import (
	"bufio"
	"literatorture/dictionary"
	"literatorture/input"
	"literatorture/messages"
	"os"
	"strings"
)

var reader bufio.Reader
var dict dictionary.Dictionary

func init() {
	reader = *bufio.NewReader(os.Stdin)
	dict = dictionary.Dictionary{}
}

func main() {

	running := true
	messages.PrintWelcomeMessage()

	for running {
		messages.PrintMenu()
		input, inputErr := input.ReadInt()
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
	word, wordErr := input.ReadStringInput(&reader)
	definition, defErr := input.ReadStringInput(&reader)

	if wordErr != nil {
		messages.PrintError(wordErr)
	} else if defErr != nil {
		messages.PrintError(defErr)
	} else {
		key := strings.ToLower(word)
		entry := dictionary.NewEntry(word, definition)
		dict.AddWord(key, *entry)
	}
}
