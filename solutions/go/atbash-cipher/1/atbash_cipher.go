package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	f := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}

		if unicode.IsLetter(r) {
			if r >= 'A' && r <= 'Z' {
				r = r + 'a' - 'A'
			}
			return 'z' - (r - 'a')
		}
		return -1
	}
	cleaned := strings.Map(f, s)
	var res strings.Builder
	res.Grow(len(cleaned) + len(cleaned)/5)

	for i, r := range cleaned {

		if i%5 == 0 && i > 0 {
			res.WriteByte(' ')
		}
		res.WriteRune(r)
	}
	return res.String()
}
