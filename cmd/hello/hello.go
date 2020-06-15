package main

import (
	"fmt"
	alias "github.com/raghuP9/golang-tutorial/pkg/triangle"
)


var a = 1
const myConst = 5
//
// add ...
//
func add(x, y int) int {
	return x+y
}

//
// swap ...
//
func swap(x, y int32) (int32, int32) {
	return y, x
}


func geometry(s int) (int, int) {
	return 4*s, s*s
}


func main() {
	perimeter := alias.Perimeter(3,4,5)
	fmt.Printf("Perimeter of triangle: %d\n", perimeter)
	areaT := alias.Area(3,4)
	fmt.Printf("Area of triangle: %v\n", areaT)

	t := alias.Triangle{
		Side1: 3,
		Side2: 4,
		Side3: 5,
	}
	perimeter = alias.PerimeterT(t)
	fmt.Printf("Perimeter: %d\n", perimeter)
	fmt.Printf("Perimeter: %d\n", t.Perimeter())
	fmt.Printf("Area: %v\n", t.Area())

}
