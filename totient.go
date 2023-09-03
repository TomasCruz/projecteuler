package projecteuler

func Totient(n int, factorSlice []int) int {
	totient := n

	for _, f := range factorSlice {
		totient /= f
		totient *= f - 1
	}

	return totient
}
