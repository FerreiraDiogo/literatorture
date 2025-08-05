package main

import (
	"bufio"
	"fmt"
	"literatorture/dictionary"
	"literatorture/fileUtils"
	"literatorture/input"
	"literatorture/messages"
	"os"
	"strings"
)

var reader bufio.Reader
var dict dictionary.Dictionary

func init() {
	reader = *bufio.NewReader(os.Stdin)
	dict = *dict.NewDictionary()
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
	case 2:
		removeWord()
		return true
	case 3:
		findWordMeaning()
		return true
	case 4:
		saveToFile()
		return true
	case 5:
		showStats()
		return true
	default:
		messages.PrintInvalidOptionMessage()
		return true
	}
}

// Reads user input and writes the content of the dict into specified file
func saveToFile() {
	if len(dict.Words) == 0 {
		fmt.Println("No words were added to your dictionary yet")
		return
	}
	fmt.Println("Insert 0 for JSON or 1 for CSV")
	input, wordErr := input.ReadInt()
	if wordErr != nil {
		messages.PrintError(wordErr)
		return
	}
	writeErr := fileUtils.Write(fileUtils.FileType(input), dict)
	if writeErr != nil {
		messages.PrintError(writeErr)
	}
}

// Shows the stats of the dict
func showStats() {
	messages.PrintStats(dict.GetStats())
}

// finds a word meaning.
func findWordMeaning() {
	messages.PrintFindWord()
	word, wordErr := input.ReadStringInput(&reader)
	if wordErr != nil {
		messages.PrintError(wordErr)
	} else {
		entry, ok := dict.Words[word]
		if ok {
			messages.PrintEntry(entry)
		} else {
			sugestions, sugestionsOk := dict.SearchOnAuxiliaryIndexes(word)
			if sugestionsOk {
				messages.PrintSugestions(sugestions)
			} else {
				messages.PrintWordDoesntExist()
			}
		}
	}
}

// Removes a given word.
func removeWord() {
	messages.PrintRemoveWordMessage()
	word, wordErr := input.ReadStringInput(&reader)
	if wordErr != nil {
		messages.PrintError(wordErr)
	} else {
		_, ok := dict.Words[word]
		if !ok {
			messages.PrintWordDoesntExist()
		} else {
			delete(dict.Words, word)
			dict.DecreaseCharacterCount(string(word[0]))
			messages.PrintRemovedMessage()
		}
	}
}

// Inserts a word and its meaning in the dictionary
// if a key for the given word is provided, then updates it
func insertWord() {
	messages.PrintInputWordMessage()
	word, wordErr := input.ReadStringInput(&reader)
	messages.PrintInputDefinitionMessage()
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
