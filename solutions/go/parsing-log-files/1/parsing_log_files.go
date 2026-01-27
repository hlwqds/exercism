package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return reg.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[~*=-]*>`)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`"(?i).*password.*"`)
	count := 0
	for _, line := range lines {
		if reg.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line\d+`)
	return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)
	result := []string{}
	for _, line := range lines {
		matches := reg.FindStringSubmatch(line)
		if len(matches) > 1 {
			userName := matches[1]
			line = fmt.Sprintf("[USR] %s %s", userName, line)
		}
		result = append(result, line)
	}
	return result
}
