package filemanipulator

import (
	"errors"
	"os"
)

func open(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.New("couldn' open file")
	}
	return file, nil
}
