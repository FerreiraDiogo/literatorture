package messages

import "fmt"

const InvalidType = "Inputed value is not a valid word"

func PrintWelcomeMessage() {
	fmt.Println("========== Literatorture ==========")
}

func PrintGoodbyeMessage() {
	fmt.Println("Thanks for using Literatorture")
}

func PrintMenu() {
	fmt.Println("==========    Options    ==========")
	fmt.Print("1 - Insert Word\n2 - Remove Word\n3 - View Word's Meaning\n4 - Update Word\n5 - Export File\n")
}

func PrintError(e error) {
	fmt.Println(e)
}

func PrintInvalidOptionMessage() {
	fmt.Println("Invalid Option")
}
