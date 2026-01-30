package romannumerals

import (
	"errors"
	"strings"
)

// 定义一个结构体来存储数值和对应的罗马字符
type romanMapping struct {
	value  int
	symbol string
}

// 按照从大到小的顺序排列，这样我们只需要一次遍历
var romanTable = []romanMapping{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid value: Roman numerals typically handle 1-3999")
	}

	var builder strings.Builder

	for _, m := range romanTable {
		for input >= m.value {
			builder.WriteString(m.symbol)
			input -= m.value
		}
	}

	return builder.String(), nil
}
