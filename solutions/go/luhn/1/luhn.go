package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	sum := 0
	if len(id) <= 1 {
		return false
	}
	idByte := []byte(id)
	for i := len(idByte) - 1; i >= 0; i-- {
		w := idByte[i]
		if !unicode.IsDigit(rune(w)) {
			return false
		}
		d := (w - '0')
		if (len(id)-i)%2 == 0 {
			d = d * 2
			if d > 9 {
				sum += int(d) - 9
			} else {
				sum += int(d)
			}
		} else {
			sum += int(d)
		}
	}
	return sum%10 == 0
}
