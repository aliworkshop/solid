package solid

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Sum() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Calculate() float64 {
	return c.Sum()
}
