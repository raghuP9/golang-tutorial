package main

import (
	"fmt"
	alias "github.com/raghuP9/golang-tutorial/pkg/triangle"
	"github.com/raghuP9/golang-tutorial/pkg/square"
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


	var sq *square.Square
	sqptr := &sq
	fmt.Printf("Square type %T\n", sq)
	sq = &square.Square{5}
	fmt.Printf("Square type %T\n", sq)
	fmt.Printf("Square: %v\n", &sq)
	sq.EnlargeSquarePtr()
	fmt.Printf("Square: %v\n", sq)
	fmt.Printf("Area: %d\n",sq.Area())
	//fmt.Printf("Area: %d\n",*sq.Area())
}

/*
var slice  []int
var slice = make([]int, 5)
var slice2 []int
slice[1:2]
slice = append(slice, slice2...)

for _, v := range slice {
   
}


var mIncpetion map[string]map[string]square.Square

var m  map[string]square.Square{
	"mystr": map[string]square.Square{
		"str": square.Square{5},
	}
	"mystr2": map[string]square.Square{
		"str2": square.Square{6},
	},
}

len(m)


if v,ok := m["mystr"]; !ok {

}

for k, _ := range m {
}
*/




