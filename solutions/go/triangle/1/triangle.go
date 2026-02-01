// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Kind .
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	sides := [3]float64{a, b, c}
	sort.Float64s(sides[:])

	for _, s := range sides {
		if s <= 0 || math.IsNaN(s) || math.IsInf(s, 0) {
			return NaT
		}
	}

	if sides[0]+sides[1] <= sides[2] {
		return NaT
	}
	switch {
	case sides[0] == sides[2]:
		return Equ
	case sides[0] == sides[1] || sides[1] == sides[2]:
		return Iso
	default:
		return Sca
	}
}
