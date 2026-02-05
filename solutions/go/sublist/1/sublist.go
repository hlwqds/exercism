package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	longList := l1
	shortList := l2
	contain := false
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	if len(l1) < len(l2) {
		longList = l2
		shortList = l1
	}
	if len(shortList) == 0 {
		contain = true
	}

	for i := 0; i <= len(longList)-len(shortList); i++ {
		for j := 0; j < len(shortList); j++ {
			if shortList[j] != longList[j+i] {
				break
			}
			if j == len(shortList)-1 {
				contain = true
				break
			}
		}
		if contain {
			break
		}
	}

	if contain && len(l1) == len(l2) {
		return RelationEqual
	}
	if contain && len(l1) < len(l2) {
		return RelationSublist
	}
	if contain && len(l1) > len(l2) {
		return RelationSuperlist
	}
	return RelationUnequal
}
