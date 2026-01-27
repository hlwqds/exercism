package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(cal FodderCalculator, num int) (float64, error) {
	amount, err := cal.FodderAmount(num)
	if err != nil {
		return 0, err
	}

	factor, err := cal.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amount * factor / float64(num), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(cal FodderCalculator, num int) (float64, error) {
	if num > 0 {
		return DivideFood(cal, num)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	num int
}

func (i InvalidCowsError) Error() string {
	desc := ""
	switch {
	case i.num < 0:
		desc = "there are no negative cows"
	case i.num == 0:
		desc = "no cows don't need food"
	}
	return fmt.Sprintf("%d cows are invalid: %s", i.num, desc)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(num int) error {
	if num <= 0 {
		return &InvalidCowsError{num}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
