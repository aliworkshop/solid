package solid

import "math"

type Cuboid struct {
	Length float64
}

func (c *Cuboid) Volume() float64 {
	return math.Pow(c.Length, 3)
}

func (c *Cuboid) Calculate() float64 {
	return c.Volume()
}
