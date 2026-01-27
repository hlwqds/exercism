package isogram

import "unicode"

func IsIsogram(word string) bool {
	wordMap := map[rune]int{}

	for _, v := range word {
		if unicode.IsLetter(v) {
			v = unicode.ToLower(v)
			if wordMap[v] != 0 {
				return false
			}
		}
		wordMap[v]++
	}
	return true
}
