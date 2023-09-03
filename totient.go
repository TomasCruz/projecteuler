package projecteuler

func Totient(n int, factorSlice []int) int {
	factorProduct := 1
	totient := 1

	for _, f := range factorSlice {
		factorProduct *= f
		totient *= f - 1
	}

	return (totient * n) / factorProduct
}

func NumDividedByTotient(n int, factorSlice []int) float64 {
	factorProduct := 1
	totient := 1

	for _, f := range factorSlice {
		factorProduct *= f
		totient *= f - 1
	}

	return float64(factorProduct) / float64(totient)
}
