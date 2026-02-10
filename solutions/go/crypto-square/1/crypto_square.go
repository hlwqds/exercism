package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	pt = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, pt)

	length := len(pt)
	if length == 0 {
		return ""
	}

	// 2. 严格遵循 c >= r 且 c-r <= 1
	c := int(math.Ceil(math.Sqrt(float64(length))))
	r := c
	if c*(c-1) >= length {
		r = c - 1
	}

	// 3. 构建结果：直接按“列”遍历
	// 我们要输出 c 条 chunk，每条长度为 r
	chunks := make([]string, c)
	for i := range c {
		var b strings.Builder
		for j := 0; j < r; j++ {
			// 矩形坐标 (j, i) 对应原字符串索引为 j*c + i
			pos := j*c + i
			if pos < length {
				b.WriteByte(pt[pos])
			} else {
				b.WriteByte(' ') // 补齐矩形缺角
			}
		}
		chunks[i] = b.String()
	}

	return strings.Join(chunks, " ")
}
