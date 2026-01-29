// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

func Hey(remark string) string {
	// 1. 处理空格：这一步是关键！
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	// 2. 判定特征
	isQuestion := strings.HasSuffix(remark, "?")

	// 判定是否在喊叫：包含大写字母且不包含任何小写字母
	hasUpper := strings.IndexFunc(remark, func(r rune) bool { return r >= 'A' && r <= 'Z' }) != -1
	hasLower := strings.IndexFunc(remark, func(r rune) bool { return r >= 'a' && r <= 'z' }) != -1
	isYelling := hasUpper && !hasLower

	// 3. 根据规则返回
	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
