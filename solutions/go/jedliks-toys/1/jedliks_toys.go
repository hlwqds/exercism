package jedlik

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain < 0 {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	for c.battery-c.batteryDrain >= 0 {
		c.Drive()
	}
	return c.distance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
