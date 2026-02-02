package prime

import (
	"errors"
	"math"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be at least 1")
	}

	// 1. 估算上限：为小规模 n 提供保底值
	var limit int
	if n < 6 {
		limit = 15 // 足够覆盖前 5 个素数 (2, 3, 5, 7, 11)
	} else {
		fn := float64(n)
		ln := math.Log(fn)
		lnln := math.Log(ln)
		limitf := fn * (ln + lnln)

		if limitf > float64(math.MaxInt) {
			return 0, errors.New("limit exceeds platform int range")
		}
		limit = int(limitf)
	}

	return sieve(n, limit)
}

func sieve(n int, limit int) (int, error) {
	// 使用 []bool，在 Go 中这是很高效的内存表达
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	count := 0
	for p := 2; p <= limit; p++ {
		if isPrime[p] {
			count++
			if count == n {
				return p, nil
			}

			// 可移植性核心：防止 p*p 溢出
			// 只有当 p*p <= limit 时才需要进行标记
			if p <= limit/p {
				for i := p * p; i <= limit; i += p {
					isPrime[i] = false
				}
			}
		}
	}
	return 0, errors.New("not found")
}
