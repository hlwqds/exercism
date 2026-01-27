package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](list []T, filter func(T) bool) []T {
	res := make([]T, 0, len(list))
	for _, v := range list {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func Discard[T any](list []T, filter func(T) bool) []T {
	return Keep(list, func(v T) bool {
		return !filter(v)
	})
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
