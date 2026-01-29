// Package pangram implements a high-performance check for pangrams.
package pangram

// AllLettersMask is a bitset where the low 26 bits are 1 (0x3FFFFFF).
// Each bit represents an encountered letter of the English alphabet.
const AllLettersMask = (1 << 26) - 1

// bitPos maps an ASCII letter to its alphabet index (0-25).
// It supports both lowercase and uppercase, returning (index, true).
// For non-alphabet characters, it returns (0, false).
func bitPos(char rune) (uint32, bool) {
	if char >= 'a' && char <= 'z' {
		return uint32(char - 'a'), true
	}
	if char >= 'A' && char <= 'Z' {
		return uint32(char - 'A'), true
	}
	return 0, false
}

// IsPangram checks if the input string uses all 26 letters of the English alphabet.
// It performs zero memory allocations and uses a bitmask for O(n) efficiency.
func IsPangram(input string) bool {
	var found uint32
	for _, char := range input {
		if pos, ok := bitPos(char); ok {
			found |= (1 << pos)
		}
		// Early exit: if all 26 bits are set, we're done.
		if found == AllLettersMask {
			return true
		}
	}
	return found == AllLettersMask
}
