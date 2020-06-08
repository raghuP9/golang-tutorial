package triangle

import "testing"


var inp = map[int][]int{
	14: []int{3,5,6},
	28: []int{8,9,10},
	9: []int{1,2,3},
}

func TestPerimeter(t *testing.T) {
	p := Perimeter(3,4,5)
	if p != 12 {
		t.Fatalf("Expected output: %d Got Output: %d", 11, p)
	}
	for k, v := range inp {
		p = Perimeter(v[0],v[1],v[2])
		if k != p{
			t.Errorf("Expected output: %d Got Output: %d", k, p)
		}
	}
}
