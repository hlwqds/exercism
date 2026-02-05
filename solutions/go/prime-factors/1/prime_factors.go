package prime

func Factors(n int64) []int64 {
	res := make([]int64, 0)
	for n%2 == 0 {
		res = append(res, 2)
		n /= 2
	}
	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}
	}
	if n > 1 {
		res = append(res, n)
	}
	return res
}
