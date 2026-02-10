package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

type MatchRes struct {
	won   int
	draw  int
	loss  int
	point int
}

type internalSort struct {
	team string
	res  MatchRes
}

func Tally(reader io.Reader, writer io.Writer) error {
	m := map[string]MatchRes{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		desc := strings.Split(line, ";")
		if len(desc) != 3 {
			return errors.New("format error")
		}

		res0, exists := m[desc[0]]
		if !exists {
			res0 = MatchRes{0, 0, 0, 0}
		}
		res1, exists := m[desc[1]]
		if !exists {
			res1 = MatchRes{0, 0, 0, 0}
		}
		switch desc[2] {
		case "win":
			res0.won++
			res1.loss++
			res0.point += 3
		case "draw":
			res0.draw++
			res1.draw++
			res0.point++
			res1.point++
		case "loss":
			res0.loss++
			res1.won++
			res1.point += 3
		default:
			return errors.New("res format")
		}
		m[desc[0]] = res0
		m[desc[1]] = res1
	}
	_, err := writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	if err != nil {
		return err
	}

	tmpSlice := make([]internalSort, 0, len(m))
	for team, res := range m {
		tmpSlice = append(tmpSlice, internalSort{team, res})
	}
	slices.SortFunc(tmpSlice, func(a, b internalSort) int {
		if b.res.point == a.res.point {
			return strings.Compare(a.team, b.team)
		}
		return b.res.point - a.res.point
	})

	for _, v := range tmpSlice {
		tmp := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", v.team, v.res.draw+v.res.loss+v.res.won, v.res.won, v.res.draw, v.res.loss, v.res.point)
		_, err := writer.Write([]byte(tmp))
		if err != nil {
			return err
		}
	}
	return nil
}
