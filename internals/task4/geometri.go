package task4

import (
	"fmt"
	"math"
)

type Geometri interface {
	Area() float64
}
type Circle struct {
	JariJari float64
}

type Rectangle struct {
	Width, Height float64
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func (ci Circle) Area() float64 {
	return math.Pi * ci.JariJari * ci.JariJari
}

func Calculator(geo Geometri) string {
	if _, isRectangle := geo.(Rectangle); isRectangle {
		return fmt.Sprintf("Rectangle: %.2f", geo.Area())
	}
	return fmt.Sprintf("Circle: %.2f", geo.Area())

}
