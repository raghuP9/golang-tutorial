package geometry

// Geometry ...
func Geometry(s int) (int, int) {
	return 4*s, square(s)
}


func square(s int) int {
	return s*s
}
