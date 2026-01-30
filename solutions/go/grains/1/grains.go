package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("invalid value")
	}

	return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	// var total uint64
	// for i := range 64 {
	// 	val, _ := Square(i + 1)
	// 	total += val
	// }
	// return total
	return math.MaxUint64
}
