package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}

	words := strings.FieldsFunc(phrase, f)

	res := make(map[string]int)
	for _, w := range words {
		w = strings.ToLower(w)
		w = strings.Trim(w, "'")
		if w != "" {
			res[w]++
		}
	}
	return res
}
