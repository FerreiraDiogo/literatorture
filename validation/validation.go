package validation

import (
	"errors"
	"literatorture/messages"
	"strconv"
)

const INVALID_CHARACTERS = "=+-_)(*&¨%$#@!\"'\\/?;:.>,<]}~^[{´}`)"

// Takes a string and validates if it is valid.
// A string is considered a valid word if it can't be converted to numeric types
// or does not contains special characters. Returns True if these conditions are
// satisfied, false otherwise
func IsValidWord(word string) (bool, error) {
	_, intErr := strconv.Atoi(word)
	if intErr == nil {
		return false, errors.New(messages.InvalidType)
	}
}
