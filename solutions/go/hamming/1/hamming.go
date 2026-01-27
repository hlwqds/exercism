package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("uneuqal len")
	}

	count := 0

	for i, v := range a {
		if v != []rune(b)[i] {
			count++
		}
	}
	return count, nil
}
