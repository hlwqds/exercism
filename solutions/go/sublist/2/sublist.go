package sublist

import "slices"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if slices.Equal(l1, l2) {
		return RelationEqual
	}

	if len(l1) < len(l2) {
		if isSublist(l1, l2) {
			return RelationSublist
		}
	}

	if len(l2) < len(l1) {
		if isSublist(l2, l1) {
			return RelationSuperlist
		}
	}
	return RelationUnequal
}

func isSublist(short, long []int) bool {
	if len(short) == 0 {
		return true
	}

	for i := 0; i <= len(long)-len(short); i++ {
		if slices.Equal(short, long[i:i+len(short)]) {
			return true
		}
	}
	return false
}
