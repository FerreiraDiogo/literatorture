package messages

import (
	"fmt"
	"literatorture/dictionary"
)

const InvalidType = "inputed value is not a valid word"

func PrintWelcomeMessage() {
	fmt.Println("========== Literatorture ==========")
}

func PrintGoodbyeMessage() {
	fmt.Println("Thanks for using Literatorture")
}

func PrintMenu() {
	fmt.Println("==========    Options    ==========")
	fmt.Print("1 - Insert/Update Word\n2 - Remove Word\n3 - View Word's Meaning\n4 - Export File\n5 - Show Stats")
}

func PrintInputWordMessage() {
	fmt.Print("\nType in the word. If it already exists, will be updated:")
}

func PrintInputDefinitionMessage() {
	fmt.Print("\nNow type in it's definition: ")
}

func PrintRemoveWordMessage() {
	fmt.Println("Type in the word you want to remove:")
}

func PrintWishToInsertMessage() {
	fmt.Println("Do you wish to insert it?")
}

func PrintWordDoesntExist() {
	fmt.Println("Inputed word doesn't exists.")
}

func PrintRemovedMessage() {
	fmt.Println("Word succesfully removed")
}

func PrintEntry(entry dictionary.Entry) {
	fmt.Printf("%s's definition is %s", entry.Word, entry.Meaning)
}

func PrintError(e error) {
	fmt.Println(e)
}

func PrintInvalidOptionMessage() {
	fmt.Println("Invalid Option")
}

func PrintSugestions(sugestions []string) {
	fmt.Println("The typed word did not foun an exact match. The sugested words are the following:")
	for _, s := range sugestions {
		fmt.Printf("- %s\n", s)
	}
}

func PrintFindWord() {
	fmt.Println("Type the word you want to search. If there's no direct match you will be offered similar words to search for.")
}

func PrintStats(stats map[string]int) {
	fmt.Println("Here is the distribution of words by initial character:")
	fmt.Println("|------------------------|")
	fmt.Printf("|  Character |  Amount   |\n")
	fmt.Println("|------------------------|")
	for i, v := range stats {
		fmt.Printf("|       %s    |     %d     |\n", i, v)
		fmt.Println("|------------------------|")
	}
}
