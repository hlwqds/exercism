package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}
	res := make([]int, 0, limit+1)
	isPrime := make([]bool, limit+1)
	for i := 0; i <= limit; i++ {
		isPrime[i] = true
	}
	for p := 2; p <= limit; p++ {
		if isPrime[p] {
			res = append(res, p)
			if p <= limit/p {
				for i := p * p; i <= limit; i += p {
					isPrime[i] = false
				}
			}

		}
	}
	return res
}
