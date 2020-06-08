package main

import (
	"fmt"
	geo "github.com/raghuP9/golang-tutorial/pkg/geometry"
	"github.com/raghuP9/golang-tutorial/pkg/triangle"
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
	x, y := int32(1), int32(2)
	x, y = swap(x,y)
	fmt.Println(x)
	fmt.Println(y)
	a = 9
	perimeter,area := geo.Geometry(myConst)
	fmt.Printf("Perimeter: %d Area: %d\n", perimeter, area)
	fmt.Println("hello-world")
	fmt.Println(add(2,5))

	perimeter = triangle.Perimeter(3,4,5)
	fmt.Printf("Perimeter of triangle: %d\n", perimeter)
}
