package bottlesong

import (
	"fmt"
)

var numberTitleMap = [11]string{
	0:  "No",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

var numberLastMap = [11]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	if takeDown > startBottles || startBottles > 10 {
		return nil
	}
	res := make([]string, 0, takeDown*5)

	for i := range takeDown {
		if i != 0 {
			res = append(res, "")
		}
		s := ""
		if startBottles != 1 {
			s = "s"
		}
		firstline := fmt.Sprintf("%s green bottle%s hanging on the wall,", numberTitleMap[startBottles], s)
		res = append(res, firstline)
		res = append(res, firstline)
		res = append(res, "And if one green bottle should accidentally fall,")
		startBottles--
		if startBottles == 1 {
			s = ""
		} else {
			s = "s"
		}
		res = append(res, fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", numberLastMap[startBottles], s))
	}
	return res
}
