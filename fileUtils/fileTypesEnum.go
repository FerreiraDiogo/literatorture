package fileUtils

import (
	"encoding/csv"
	"encoding/json"
	"literatorture/dictionary"
	"log"
	"os"
)

type FileType int

const (
	JSON FileType = iota
	CSV
)

var filename = map[FileType]string{
	JSON: "dict.json",
	CSV:  "dict.csv",
}

var writeToFile = map[FileType]func(*os.File, dictionary.Dictionary) error{

	JSON: writeToJson,
	CSV:  writeToCsv,
}

func (f FileType) String() string {
	return filename[f]
}

func (f FileType) writeToFile() func(*os.File, dictionary.Dictionary) error {
	return writeToFile[f]
}

func writeToJson(file *os.File, dict dictionary.Dictionary) error {
	log.Printf("Writing JSON file\n")
	jsonified, marshErr := json.MarshalIndent(dict.Words, "", " ")
	if marshErr != nil {
		log.Printf("An error has occurred during JSON marshaling: %s\n", marshErr)
		return marshErr
	}
	file.Write(jsonified)
	log.Println("JSON file successfully writen")
	return nil
}
func writeToCsv(file *os.File, dict dictionary.Dictionary) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := [][]string{
		{"Word", "Meaning"},
	}

	for _, v := range dict.Words {
		records = append(records, []string{v.Word, v.Meaning})
	}

	writer.WriteAll(records)

	return nil
}
