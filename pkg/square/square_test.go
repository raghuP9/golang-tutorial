package square

import "testing"

var inp = map[int][]int{
	2: []int{8, 4},
	5: []int{20, 25},
	9: []int{36, 81},
}


func TestPerimeter(t *testing.T) {
	for k, v := range inp {
		sq := Square{k}
		 
		if p := sq.Perimeter(); p != v[0] {
			t.Errorf("Expected: %d Output: %d\n", v[0], p)
		}
		if sq.Area() != v[1] {
			t.Errorf("Expected: %d Output: %d\n", v[0], sq.Area())
		}
	}
}
