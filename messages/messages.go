package messages

import "fmt"

func PrintMenu() {
	fmt.Println("========== Literatorture ==========")
	fmt.Println("==========    Options    ==========")
	fmt.Print("1 - Insert Word\n2 - Remove Word\n3 - View Word's Meaning\n4 - Update Word\n5 - Export File\n")

}
