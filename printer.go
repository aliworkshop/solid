package solid

import (
	"encoding/json"
	"fmt"
)

type shape struct {
	object ManageShape
}

func NewShape(s ManageShape) Printer {
	return &shape{object: s}
}

func (s *shape) Print() {
	fmt.Println(fmt.Sprintf("(area sum/volume) of this %T is %f", s.object, s.object.Calculate()))
}

func (s *shape) Json() {
	b, err := json.Marshal(map[string]interface{}{
		fmt.Sprintf("%T", s.object): s.object.Calculate(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
