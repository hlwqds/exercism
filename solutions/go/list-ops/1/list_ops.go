package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return s
	}
	number := 0
	for i := range s {
		if fn(s[i]) {
			number++
		}
	}
	if number == 0 {
		return nil
	}

	res := make(IntList, number)
	number = 0
	for i := range s {
		if fn(s[i]) {
			res[number] = s[i]
			number++
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, len(s))
	for i, v := range s {
		res[i] = fn(v)
	}
	return res
}

func (s IntList) Reverse() IntList {
	res := make(IntList, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	if len(lst) == 0 {
		return s
	}
	res := make(IntList, len(s)+len(lst))
	copy(res, s)
	copy(res[len(s):], lst)
	return res
}

func (s IntList) Concat(lists []IntList) IntList {
	length := 0
	for i := range lists {
		length += len(lists[i])
	}
	if length == 0 {
		return s
	}
	length += len(s)
	if length == 0 {
		return nil
	}

	res := make(IntList, length)
	copy(res, s)
	copied := len(s)
	for i := range lists {
		copy(res[copied:], lists[i])
		copied += len(lists[i])
	}

	return res
}
