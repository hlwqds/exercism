// Package reverse implements a reverse func for reversing words in string
package reverse

// isPureASCII check if s is pure ASCII
func isPureASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 128 {
			return false
		}
	}
	return true
}

func reverseInPlace[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseRune(input string) string {
	r := []rune(input)
	reverseInPlace(r)
	return string(r)
}

func reverseASCII(input string) string {
	b := []byte(input)
	reverseInPlace(b)
	return string(b)
}

// Reverse returns a new string with the characters of input in reverse order.
func Reverse(input string) string {
	if len(input) <= 1 {
		return input
	}
	if isPureASCII(input) {
		return reverseASCII(input)
	}
	return reverseRune(input)
}
