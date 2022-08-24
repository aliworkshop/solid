package solid

type Shape interface {
	Sum() float64
}

type ThreeDimensionalShape interface {
	Volume() float64
}

type ManageShape interface {
	Calculate() float64
}

type Printer interface {
	Print()
	Json()
}
