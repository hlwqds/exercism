package allyourbase

import (
	"errors"
	"slices"
)

func pow(x int, y int) int {
	res := 1
	for range y {
		res *= x
	}
	return res
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	decRes := 0
	if inputBase <= 1 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase <= 1 {
		return nil, errors.New("output base must be >= 2")
	}

	for i := len(inputDigits) - 1; i >= 0; i-- {
		if inputDigits[i] >= inputBase || inputDigits[i] < 0 {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decRes += inputDigits[i] * pow(inputBase, len(inputDigits)-1-i)
	}
	if decRes == 0 {
		return []int{0}, nil
	}

	res := make([]int, 0, len(inputDigits))
	for decRes > 0 {
		mod := decRes % outputBase
		res = append(res, mod)
		decRes /= outputBase
	}

	slices.Reverse(res)

	return res, nil
}
