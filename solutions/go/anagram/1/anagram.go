package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	lSubject := strings.ToLower(subject)
	res := make([]string, 0, len(candidates))
	r := []rune(lSubject)
	slices.Sort(r)
	for _, candidate := range candidates {
		lCandidate := strings.ToLower(candidate)
		if lCandidate == lSubject || len(lSubject) != len(lCandidate) {
			continue
		}
		rCandidate := []rune(lCandidate)
		slices.Sort(rCandidate)
		if slices.Equal(r, rCandidate) {
			res = append(res, candidate)
		}
	}
	return res
}
