package isbn

func IsValidISBN(isbn string) bool {
	count := 0
	res := 0

	for _, value := range isbn {
		if value == '-' {
			continue
		}
		if count >= 10 {
			return false
		}
		switch {
		case value >= '0' && value <= '9':
			res += int(value-'0') * (10 - count)
		case value == 'X' && count == 9:
			res += 10
		default:
			return false
		}
		count++
	}

	return res%11 == 0 && count == 10
}
