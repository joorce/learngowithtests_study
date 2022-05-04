package structs

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
type Circle struct {
	Radius float64
}
