package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var ratio float64

	// Exercism Space Age 要求的精确比率
	switch planet {
	case "Mercury":
		ratio = 0.2408467
	case "Venus":
		ratio = 0.61519726
	case "Earth":
		ratio = 1.0
	case "Mars":
		ratio = 1.8808158
	case "Jupiter":
		ratio = 11.862615
	case "Saturn":
		ratio = 29.447498
	case "Uranus":
		ratio = 84.016846
	case "Neptune":
		ratio = 164.79132
	default:
		// 如果行星不存在，通常返回 -1 或根据测试要求处理
		return -1
	}

	// 1. 使用括号确保先计算分母，提高微小的浮点数精度
	// 2. 31557600 是地球公转周期的精确秒数 (365.25天)
	return seconds / (ratio * 31557600)
}
