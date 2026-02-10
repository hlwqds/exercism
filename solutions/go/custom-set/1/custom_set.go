package stringset

import (
	"strconv"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	m map[string]struct{}
}

func New() Set {
	return Set{make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	set := Set{make(map[string]struct{}, len(l))}
	for _, v := range l {
		set.Add(v)
	}
	return set
}

func (s Set) String() string {
	res := strings.Builder{}
	res.Grow(1024)
	res.WriteByte('{')
	first := true
	for k := range s.m {
		if !first {
			res.WriteString(", ")
		}
		res.WriteString(strconv.Quote(k))
		first = false
	}
	res.WriteByte('}')
	return res.String()
}

func (s Set) IsEmpty() bool {
	return len(s.m) == 0
}

func (s Set) Has(elem string) bool {
	_, exists := s.m[elem]
	return exists
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		s.m[elem] = struct{}{}
	}
}

func Subset(s1, s2 Set) bool {
	for k := range s1.m {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s2.m {
		if s1.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.m) != len(s2.m) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	if len(s1.m) > len(s2.m) {
		s1, s2 = s2, s1
	}
	newSet := New()
	for k := range s1.m {
		if s2.Has(k) {
			newSet.Add(k)
		}
	}
	return newSet
}

func Difference(s1, s2 Set) Set {
	newSet := New()
	for k := range s1.m {
		if !s2.Has(k) {
			newSet.Add(k)
		}
	}
	return newSet
}

func Union(s1, s2 Set) Set {
	newSet := New()
	for k := range s2.m {
		newSet.Add(k)
	}
	for k := range s1.m {
		newSet.Add(k)
	}
	return newSet
}
