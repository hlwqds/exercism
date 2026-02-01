package clock

import "fmt"

// Clock type here.
type Clock struct {
	m int
}

const (
	dayMinutes = 24 * 60
)

func New(h, m int) Clock {
	total := (h*60 + m) % dayMinutes
	if total < 0 {
		total += dayMinutes
	}

	return Clock{total}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
