package triangle

import "math"

type Triangle struct {
	Side1 int
	Side2 int
	Side3 int
}

// Perimeter ...
func Perimeter(x,y,z int) int {
	return x+y+z
}

// Area ..
func Area(base, height int) float32 {
	return (float32(base)*float32(height))/2
}

func PerimeterT(t Triangle) int {
	return t.Side1 + t.Side2 + t.Side3
}
/*
func NewTriangle(x,y,z int) Triangle {
	return Triangle{
		side1:x,
		side2:y,
		side3:z,
	}
}*/

func (t Triangle) Perimeter() int {
	return t.Side1 + t.Side2 + t.Side3
}

//s = p/2
//sqrt(s*(s-a)*(s-b)*(s-c))
func (t Triangle) Area() float32 {
	p := float64(t.Perimeter())/2
	area := math.Sqrt(p*(p - float64(t.Side1))*(p - float64(t.Side2))*(p - float64(t.Side3)))
	return float32(area)
}
