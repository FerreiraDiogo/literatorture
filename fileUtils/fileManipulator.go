package fileUtils

import (
	"errors"
	"literatorture/dictionary"
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func open(fileName string) (*os.File, error) {
	logger.Info("Opening file: ", "fileName", fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.New("couldn't open file")
	}
	return file, nil
}

func Write(fileType FileType, dict dictionary.Dictionary) error {
	file, fileError := open(fileType.String())
	if fileError != nil {
		logger.Error("an error occurred during the file opening: ", "fileError", fileError)
		return fileError
	}
	defer file.Close()

	return fileType.writeToFile()(file, dict)
}
