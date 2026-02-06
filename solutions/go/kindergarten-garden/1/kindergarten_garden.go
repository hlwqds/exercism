package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

// Define the Garden type here.

type record struct {
	row int
	cow int
}

type Garden struct {
	diagram [2]string
	records map[string]int
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := Garden{[2]string{}, map[string]int{}}
	childrenTmp := make([]string, len(children))
	copy(childrenTmp, children)
	slices.Sort(childrenTmp)
	children = childrenTmp
	if diagram[0] != '\n' {
		return nil, errors.New("format")
	}
	sDiagram := strings.Split(strings.TrimSpace(diagram), "\n")
	for _, v := range sDiagram {
		for _, jv := range v {
			if jv != 'G' && jv != 'C' && jv != 'R' && jv != 'V' {
				return nil, errors.New("invalid diagram 1")
			}
		}
	}
	if len(sDiagram) != 2 {
		return nil, errors.New("invalid diagram 2")
	}
	if len(sDiagram[0]) != len(sDiagram[1]) || len(sDiagram[0]) < len(children)*2 || len(sDiagram[0])%2 != 0 {
		return nil, errors.New("cap not enough")
	}
	garden.diagram = [2]string(sDiagram)
	row := 0
	for _, child := range children {
		_, exist := garden.records[child]
		if exist {
			return nil, errors.New("duplicate")
		}
		garden.records[child] = row
		row += 2
	}
	return &garden, nil
}

func transPlant(r byte) string {
	switch r {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}

func (g *Garden) Plants(child string) ([]string, bool) {
	startIndex, exist := g.records[child]
	if !exist {
		return nil, false
	}
	res := make([]string, 0, 4)
	for i := range 2 {
		for j := startIndex; j < startIndex+2; j++ {
			res = append(res, transPlant(g.diagram[i][j]))
		}
	}
	return res, true
}
