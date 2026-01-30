package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int, 26)
	for k, v := range in {
		for _, s := range v {
			res[strings.ToLower(s)] = k
		}
	}
	return res
}
