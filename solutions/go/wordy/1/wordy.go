package wordy

import (
	"regexp"
	"strconv"
)

var precedence = map[string]int{
	"plus":       1,
	"multiplied": 1,
	"minus":      1,
	"divided":    1,
}

func Answer(question string) (int, bool) {
	reg := regexp.MustCompile(`(?i)-?\d+|[a-z]+`)
	tokens := reg.FindAllString(question, -1)
	digits := make([]int, 0, len(tokens))
	opcode := make([]string, 0, len(tokens))
	if len(tokens) <= 1 {
		return 0, false
	}

	supportedOps := map[string]bool{
		"plus": true, "minus": true, "multiplied": true, "divided": true,
	}

	for _, t := range tokens {
		if _, err := strconv.Atoi(t); err == nil {
			continue
		}
		if t == "What" || t == "is" || t == "by" {
			continue
		}

		if !supportedOps[t] {
			return 0, false
		}
	}

	expectedDigit := true
	popAndCalc := func() bool {
		if len(digits) >= 2 {

			b := digits[len(digits)-1]
			a := digits[len(digits)-2]
			res := 0
			switch opcode[len(opcode)-1] {
			case "plus":
				res = a + b
			case "minus":
				res = a - b
			case "multiplied":
				res = a * b
			case "divided":
				if b == 0 {
					return false
				}
				res = a / b
			}
			digits[len(digits)-2] = res
			digits = digits[:len(digits)-1]
			opcode = opcode[:len(opcode)-1]
			return true
		}
		return false
	}
	for _, v := range tokens {
		if v == "What" || v == "is" || v == "by" {
			continue
		}

		if val, err := strconv.Atoi(v); err == nil {
			if !expectedDigit {
				return 0, false
			}
			digits = append(digits, val)
			expectedDigit = false
		} else {
			if expectedDigit {
				return 0, false
			}
			for len(opcode) > 0 && (precedence[opcode[len(opcode)-1]] >= precedence[v]) {
				if !popAndCalc() {
					return 0, false
				}
			}
			opcode = append(opcode, v)
			expectedDigit = true
		}
	}
	for len(opcode) > 0 {
		if !popAndCalc() {
			return 0, false
		}
	}
	if len(digits) != 1 {
		return 0, false
	}
	return digits[0], true
}
