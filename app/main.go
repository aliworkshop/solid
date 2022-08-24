package main

import (
	"solid"
)

func main() {
	var shapes []solid.ManageShape
	shapes = append(shapes, &solid.Circle{Radius: 5})
	shapes = append(shapes, &solid.Square{Length: 6})
	shapes = append(shapes, &solid.Cuboid{Length: 4})

	for _, shape := range shapes {
		solid.NewShape(shape).Print()
	}
}
