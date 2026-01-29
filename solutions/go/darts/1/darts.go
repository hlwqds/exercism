package darts

const (
	outerRadiusSq  = 100.0
	middleRadiusSq = 25.0
	innerRadiusSq  = 1.0
)

func Score(x, y float64) int {
	distSq := x*x + y*y
	switch {
	case distSq > outerRadiusSq:
		return 0
	case distSq > middleRadiusSq:
		return 1
	case distSq > innerRadiusSq:
		return 5
	default:
		return 10
	}
}
