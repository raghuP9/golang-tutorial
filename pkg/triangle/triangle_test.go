package triangle

import "testing"


var inp = map[int][]int{
	14: []int{3,5,6},
	27: []int{8,9,10},
	6: []int{1,2,3},
}

var inp_1 = map[float32][]int{
	9: []int{3,6},
	36: []int{8,9},
	1: []int{1,2},
}

func TestPerimeter(t *testing.T) {
	p := Perimeter(3,4,5)
	if p != 12 {
		t.Errorf("Expected output: %d Got Output: %d", 11, p)
	}
	for k, v := range inp {
		p = Perimeter(v[0],v[1],v[2])
		if k != p{
			t.Errorf("Expected output: %d Got Output: %d", k, p)
		}
		tri := Triangle{v[0],v[1],v[2]}
		p = tri.Perimeter()

		if k != p{
			t.Errorf("Expected output: %d Got Output: %d", k, p)
		}
	}
}

func TestArea(t *testing.T) {
	for k, v := range inp_1 {
		p := Area(v[0],v[1])
		if k != p{
			t.Errorf("Expected output: %v Got Output: %v", k, p)
		}
	}
}
