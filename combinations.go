package projecteuler

// Combinations calculates and returns combinations of k out of n elements. n has to be greater or equal to k.
// Each slice in the resulting matrix has 1s per elements used in a particular combination.
func Combinations(n, k byte, f func(...interface{}) bool, args ...interface{}) (
	combinations [][]byte, retValue bool) {

	if k == 0 {
		//combinations = [][]byte{[]byte{}}
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
	for _, c := range combWO {
		newC := append(c, 0)
		combinations = append(combinations, newC)
		retValue = processComb(newC, f, args...)
		if retValue {
			return
		}
	}

	// combinations with the last element
	combW, _ := Combinations(n-1, k-1, nil)
	for _, c := range combW {
		newC := append(c, 1)
		combinations = append(combinations, newC)
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
