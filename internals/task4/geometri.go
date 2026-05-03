package task4

import (
	"fmt"
	"math"
)

type Geometri interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Calculator(total []Geometri) string {
	var result float64
	for _, shape := range total {
		result += shape.Area()
	}
	return fmt.Sprintf("total area: %v", result)
}

func Luas(hasil Geometri) string {
	luas := hasil.Area()
	return fmt.Sprintf(": %v", luas)
}
