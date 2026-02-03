package phonenumber

import (
	"errors"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, phoneNumber)

	if len(phoneNumber) != 10 && len(phoneNumber) != 11 {
		return "", errors.New("invalid number")
	}

	if phoneNumber[0] == '1' && len(phoneNumber) == 11 {
		phoneNumber = phoneNumber[1:]
	} else if len(phoneNumber) == 11 {
		return "", errors.New("invalid number1")
	}
	if phoneNumber[0] == '0' || phoneNumber[0] == '1' || phoneNumber[3] == '0' || phoneNumber[3] == '1' {
		return "", errors.New("invalid number2")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + phoneNumber[0:3] + ") " + phoneNumber[3:6] + "-" + phoneNumber[6:10], nil
}
