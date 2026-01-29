package pangram

const AllLettersMask = (1 << 26) - 1

func IsPangram(input string) bool {
	var found uint32
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			found |= (1 << (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			found |= (1 << (char - 'A'))
		}

		// 提前退出优化：一旦集齐 26 个字母，立即返回 true
		if found == AllLettersMask {
			return true
		}
	}
	return found == AllLettersMask
}
