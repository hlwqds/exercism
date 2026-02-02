package armstrong

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}
	res := 0
	count := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		count++
	}
	for tmp := n; tmp > 0; tmp /= 10 {
		digit := tmp % 10
		p := 1
		for i := 0; i < count; i++ {
			p *= digit
		}
		res += p
	}
	return res == n
}
