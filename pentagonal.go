package projecteuler

// IsPentagonal returns if x is pentagonal
func IsPentagonal(x int, roots map[int]int) bool {
	squared := 24*x + 1
	if root, ok := roots[squared]; ok {
		return root%6 == 5
	}

	return false
}

// IsTriangular returns if x is triangular
func IsTriangular(x int, roots map[int]int) bool {
	squared := 8*x + 1
	if root, ok := roots[squared]; ok {
		return root%2 == 1
	}

	return false
}

// IsHexagonal returns if x is hexagonal
func IsHexagonal(x int, roots map[int]int) bool {
	squared := 8*x + 1
	if root, ok := roots[squared]; ok {
		return root%4 == 3
	}

	return false
}
