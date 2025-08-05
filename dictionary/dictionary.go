package dictionary

import (
	"slices"
	"strings"
)

type Dictionary struct {
	Words             map[string]Entry
	Prefixes, Sufixes map[string][]string
	stats             map[string]int
}

// Adds a word into the dictionary
func (d *Dictionary) AddWord(key string, entry Entry) {
	d.Words[key] = entry
	d.addPrefix(key)
	d.addSufix(key)
	d.increaseCharacterCount(string(key[0]))
}

// Constructor method
func (d Dictionary) NewDictionary() *Dictionary {
	return &Dictionary{Words: make(map[string]Entry, 0), Prefixes: make(map[string][]string), Sufixes: make(map[string][]string), stats: make(map[string]int)}
}

// Adds prefixes for a given word in the Prefix inverted index
func (d *Dictionary) addPrefix(word string) {
	word = strings.TrimSpace(word)

	key := ""
	for i := 0; i < len(word); i++ {

		key += string(word[i])
		d.Prefixes[key] = append(d.Prefixes[key], word)

	}

}

// Adds sufixes for a given word in the Sufix inverted index
func (d *Dictionary) addSufix(word string) {
	word = strings.TrimSpace(word)
	invertedWord := reverse(word)
	key := ""
	for i := 0; i < len(invertedWord); i++ {
		key = string(invertedWord[i]) + key
		d.Sufixes[key] = append(d.Sufixes[key], word)
	}
}

// Reverses a string
func reverse(word string) string {
	result := ""
	for _, v := range word {
		result = string(v) + result
	}
	return result
}

// Searchs a word by prefixes or sufixes
func (d *Dictionary) SearchOnAuxiliaryIndexes(word string) ([]string, bool) {
	suggestedWords, prefixOk := d.Prefixes[word]
	if prefixOk {
		return suggestedWords, prefixOk
	} else {
		suggestedWords, sufixedOk := d.Sufixes[word]
		if sufixedOk {
			return suggestedWords, sufixedOk
		}
	}
	return make([]string, 0), false
}

// Increase the char count for the given char in the stats map
func (d *Dictionary) increaseCharacterCount(char string) {
	d.stats[char]++
}

// Decrease the char count for the given char in the stats map
func (d *Dictionary) DecreaseCharacterCount(char string) {
	d.stats[char]--
}

// Gets basic statistics for thw words in the dict
func (d *Dictionary) GetStats() map[string]int {
	keys := make([]string, 0)
	orderedMap := make(map[string]int, 0)

	for k := range d.stats {
		keys = append(keys, string(k))
	}

	slices.Sort(keys)
	for _, v := range keys {
		orderedMap[v] = d.stats[v]
	}
	return orderedMap
}
