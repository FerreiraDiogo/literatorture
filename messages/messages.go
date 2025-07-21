package messages

import "fmt"

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

func PrintError(e error) {
	fmt.Println(e)
}

func PrintInvalidOptionMessage() {
	fmt.Println("Invalid Option")
}
