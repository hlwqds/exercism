// Package twofer is a solution
package twofer

import "fmt"

// ShareWith hould have a comment documenting it.
func ShareWith(name string) string {
	title := "you"
	if name != "" {
		title = name
	}

	return fmt.Sprintf("One for %s, one for me.", title)
}
