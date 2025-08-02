package dictionary

import (
	"strings"
)

type Dictionary struct {
	Words             map[string]Entry
	Prefixes, Sufixes map[string][]string
}

func (d *Dictionary) AddWord(key string, entry Entry) {
	d.Words[key] = entry
	d.addPrefix(key)
	d.addSufix(key)
}

func (d Dictionary) NewDictionary() *Dictionary {
	return &Dictionary{Words: make(map[string]Entry, 0), Prefixes: make(map[string][]string), Sufixes: make(map[string][]string)}
}

func (d *Dictionary) addPrefix(word string) {
	word = strings.TrimSpace(word)

	key := ""
	for i := 0; i < len(word); i++ {

		key += string(word[i])
		d.Prefixes[key] = append(d.Prefixes[key], word)

	}

}
func (d *Dictionary) addSufix(word string) {
	word = strings.TrimSpace(word)
	invertedWord := reverse(word)
	key := ""
	for i := 0; i < len(invertedWord); i++ {
		key = string(invertedWord[i]) + key
		d.Sufixes[key] = append(d.Sufixes[key], word)
	}
}

func reverse(word string) string {
	result := ""
	for _, v := range word {
		result = string(v) + result
	}
	return result
}

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
