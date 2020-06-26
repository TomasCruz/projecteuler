package projecteuler

// Combinations calculates and returns combinations of k out of n elements. n has to be greater or equal to k.
// Each slice in the resulting matrix has 1s per elements used in a particular combination.
func Combinations(n, k byte, f func(...interface{}) bool, args ...interface{}) (
	combinations [][]byte, retValue bool) {

	if k == 0 {
		slZeroes := make([]byte, n)
		combinations = [][]byte{slZeroes}
		return
	}

	if n == k {
		slOnes := make([]byte, n)
		for i := 0; i < int(n); i++ {
			slOnes[i] = 1
		}
		combinations = [][]byte{slOnes}
		retValue = processComb(slOnes, f, args...)
		return
	}

	// combinations without the last element
	combWO, _ := Combinations(n-1, k, nil)
	combinations, retValue = appendLastElement(0, combWO, f, args...)
	if retValue {
		return
	}

	// combinations with the last element
	combW, _ := Combinations(n-1, k-1, nil)
	combW, retValue = appendLastElement(1, combW, f, args...)
	if retValue {
		return
	}
	combinations = append(combinations, combW...)

	return
}

func appendLastElement(element byte, combinations [][]byte, f func(...interface{}) bool, args ...interface{}) (
	newCombinations [][]byte, retValue bool) {

	for _, c := range combinations {
		newC := append(c, element)
		newCombinations = append(newCombinations, newC)
		retValue = processComb(newC, f, args...)
		if retValue {
			return
		}
	}

	return
}

func processComb(currComb []byte, f func(...interface{}) bool, args ...interface{}) (retValue bool) {
	if f != nil {
		args = append(args, currComb)
		if f(args...) {
			retValue = true
		}
		args = args[:len(args)-1]
	}

	return
}
