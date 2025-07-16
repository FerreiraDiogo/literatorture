package validation

import (
	"errors"
	"literatorture/messages"
	"strconv"
	"strings"
)

const INVALID_CHARACTERS = "=+-_)(*&¨%$#@!\"'\\/?;:.>,<]}~^[{´}`)"

// Takes a string and validates if it is valid.
// A string is considered a valid word if it can't be converted to numeric types
// or does not contains special characters. Returns True if these conditions are
// satisfied, false otherwise
func IsValidWord(word string) error {
	_, intErr := strconv.Atoi(word)
	if intErr == nil {
		return errors.New(messages.InvalidType)
	}
	_, floatErr := strconv.ParseFloat(word, 64)
	if floatErr == nil {
		return errors.New(messages.InvalidType)
	}
	if strings.ContainsAny(word, INVALID_CHARACTERS) {
		return errors.New(messages.InvalidType)
	}
	return nil
}
