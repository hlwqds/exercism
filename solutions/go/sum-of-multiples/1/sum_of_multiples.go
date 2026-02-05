package summultiples

func SumMultiples(limit int, divisors ...int) int {
	set := map[int]bool{}
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i += divisor {
			set[i] = true
		}
	}

	res := 0

	for k := range set {
		res += k
	}
	return res
}
