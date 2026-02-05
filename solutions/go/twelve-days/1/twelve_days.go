package twelve

import "strings"

var lifts = [13]string{
	0:  "",
	1:  "a Partridge",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

var dayth = [13]string{
	0:  "",
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

const (
	format1 = "On the "
	format2 = " day of Christmas my true love gave to me: "
	format3 = " in a Pear Tree."
	format4 = ", "
	format5 = "and "
)

func Verse(i int) string {
	length := len(format1) + len(format2) + len(format3)
	length += len(dayth[i])
	for j := i; j > 0; j-- {
		if j == 1 && i >= 2 {
			length += len(format5)
		}

		length += len(lifts[j])
		if j >= 2 {
			length += len(format4)
		}
	}
	res := strings.Builder{}
	res.Grow(length)
	res.WriteString(format1)
	res.WriteString(dayth[i])
	res.WriteString(format2)
	for j := i; j > 0; j-- {
		if j == 1 && i >= 2 {
			res.WriteString(format5)
		}
		res.WriteString(lifts[j])
		if j >= 2 {
			res.WriteString(format4)
		}
	}
	res.WriteString(format3)
	return res.String()
}

func Song() string {
	var res strings.Builder
	// 12天歌词总长度大约在 2500 字节左右
	// 如果想继续硬核计算，可以像 Verse 一样算一遍，
	// 或者直接给个保守值，让它在整个程序生命周期内只分配一次
	res.Grow(2560)

	for i := 1; i <= 12; i++ {
		if i != 1 {
			res.WriteByte('\n')
		}
		res.WriteString(Verse(i))
	}
	return res.String()
}
