package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
