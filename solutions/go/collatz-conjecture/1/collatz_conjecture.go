// Package collatzconjecture for fun
package collatzconjecture

import (
	"errors"
)

var meet = map[int]bool{}

// CollatzConjecture for fun
func CollatzConjecture(n int) (int, error) {
	if n == 1 {
		clear(meet)
		return 0, nil
	}
	result := 0
	if n%2 == 0 {
		result = n / 2
	} else {
		result = n*3 + 1
	}
	if meet[result] {
		return 0, errors.New("infinity")
	}
	meet[result] = true
	step, err := CollatzConjecture(result)
	if err != nil {
		return step, err
	}
	return step + 1, nil
}
