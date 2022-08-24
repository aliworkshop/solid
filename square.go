package solid

import "math"

type Square struct {
	Length float64
}

func (s *Square) Sum() float64 {
	return math.Pow(s.Length, 2)
}

func (s *Square) Calculate() float64 {
	return s.Sum()
}
