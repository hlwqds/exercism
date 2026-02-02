package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("invalid span")
	}
	if span == 0 {
		return 1, nil
	}
	f := func(s_digits string) (int64, error) {
		var total int64 = 1
		for _, digit := range s_digits {
			if digit < '0' || digit > '9' {
				return 0, errors.New("not number")
			}
			total *= int64(digit - '0')
		}
		return int64(total), nil
	}
	var max int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		res, err := f(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		if max < res {
			max = res
		}
	}
	return max, nil
}
