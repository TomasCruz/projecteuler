package projecteuler

// ModuloSum calculates sum module mod
func ModuloSum(a, b int64, mod int64) (sum int64) {
	sum = (a + b) % mod
	return
}

// ModuloMultiply calculates multiplication module mod
func ModuloMultiply(a, b int64, mod int64) (result int64) {
	result = (a * b) % mod
	return
}

// ModuloSelfPower calculates self power of x module mod
func ModuloSelfPower(x int, mod int64) (selfPower int64) {
	selfPower = 1

	for i := 1; i <= x; i++ {
		selfPower = ModuloMultiply(selfPower, int64(x), mod)
	}

	return
}
