package square

// Square represents a geometric square object
type Square struct {
	Side int
}

// Perimeter returns perimeter of the square
func (sq Square) Perimeter() int {
	return 4*sq.Side
}

// Area returns area of the square
func (sq Square) Area() int {
	return sq.Side*sq.Side
}

// EnlargeSquare ...
func (sq Square) EnlargeSquare() {
	sq.Side = 2*sq.Side
}

// EnlargeSquarePtr ...
func (sq *Square) EnlargeSquarePtr(){
	sq.Side = 2*sq.Side
}
