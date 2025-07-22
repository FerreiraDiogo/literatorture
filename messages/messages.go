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
	fmt.Print("1 - Insert/Update Word\n2 - Remove Word\n3 - View Word's Meaning\n4 - Export File\n")
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
