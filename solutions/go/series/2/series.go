package series

func All(n int, s string) []string {
	number := len(s) - n + 1
	if number <= 0 {
		return nil
	}
	res := make([]string, 0, number)
	for i := range number {
		res = append(res, s[i:i+n])
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
