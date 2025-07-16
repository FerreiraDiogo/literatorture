package dictionary

type Dictionary struct {
	Words map[string]Entry
}

func (d *Dictionary) AddWord(key string, entry Entry) {
	d.Words[key] = entry
}
