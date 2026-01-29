// Package pangram provides utilities for identigying pangrams,
// which are sentences using every letter of a given alphabet.
package pangram

// AllLettersMask represents a bitset where the first 26 bits are set to 1.
// Binary: 00000011 11111111 11111111 11111111
const AllLettersMask = (1 << 26) - 1

func bitPos(char rune) (uint32, bool) {
	if char >= 'a' && char <= 'z' {
		return uint32(char - 'a'), true
	}
	if char >= 'A' && char <= 'Z' {
		return uint32(char - 'A'), true
	}
	return 0, false
}

func IsPangram(input string) bool {
	var found uint32
	for _, char := range input {
		if pos, ok := bitPos(char); ok {
			found |= (1 << pos)
		}
		// 提前退出优化：一旦集齐 26 个字母，立即返回 true
		if found == AllLettersMask {
			return true
		}
	}
	return found == AllLettersMask
}
