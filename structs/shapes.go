package structs

import "math"

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Shape interface {
	Area() float64
}
