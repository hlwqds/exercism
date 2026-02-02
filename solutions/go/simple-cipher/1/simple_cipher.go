package cipher

import (
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}

type vigenere struct {
	distances []int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance >= 26 || distance <= -26 {
		return nil
	}
	return shift{(distance%26 + 26) % 26}
}

func encode(input string, distance int) string {
	encodeStr := strings.Builder{}
	encodeStr.Grow(len(input))
	distance = (distance%26 + 26) % 26
	for _, r := range input {
		switch {
		case r >= 'a' && r <= 'z':
			encodeStr.WriteRune('a' + (r-'a'+rune(distance))%26)
		case r >= 'A' && r <= 'Z':
			encodeStr.WriteRune('a' + (r-'A'+rune(distance))%26)
		}
	}
	return encodeStr.String()
}

func (c shift) Encode(input string) string {
	return encode(input, c.distance)
}

func (c shift) Decode(input string) string {
	return encode(input, c.distance*-1)
}

func NewVigenere(key string) Cipher {
	length := len(key)
	haveShift := false
	distances := make([]int, length)
	for i, k := range key {
		if k < 'a' || k > 'z' {
			return nil
		}
		distances[i] = int(k - rune('a'))
		if distances[i] != 0 {
			haveShift = true
		}
	}
	if !haveShift {
		return nil
	}
	return vigenere{distances}
}

func (v vigenere) Encode(input string) string {
	length := len(v.distances)
	encodeStr := strings.Builder{}
	i := 0
	for _, r := range input {
		distance := v.distances[i%length]
		switch {
		case r >= 'a' && r <= 'z':
			encodeStr.WriteRune('a' + (r-'a'+rune(distance))%26)
			i++
		case r >= 'A' && r <= 'Z':
			encodeStr.WriteRune('a' + (r-'A'+rune(distance))%26)
			i++
		}
	}
	return encodeStr.String()
}

func (v vigenere) Decode(input string) string {
	length := len(v.distances)
	input = strings.ReplaceAll(input, " ", "")
	encodeStr := strings.Builder{}
	i := 0
	for _, r := range input {
		distance := (26 - v.distances[i%length])
		switch {
		case r >= 'a' && r <= 'z':
			encodeStr.WriteRune('a' + (r-'a'+rune(distance))%26)
			i++
		case r >= 'A' && r <= 'Z':
			encodeStr.WriteRune('a' + (r-'A'+rune(distance))%26)
			i++
		}
	}
	return encodeStr.String()
}
