package main

import (
	"math"

	"github.com/TomasCruz/projecteuler"
)

// relationship of "is a smaller prime contatenable in any order with this prime"
// will henceforth be called "kid brother". This prime might be referred as big bro
type kidBrothers struct {
	primes       []int
	primeSet     map[int]struct{}
	primeCount   int
	setSize      int
	limit        int
	winnerSum    int
	kidBros      []map[int]struct{} // per prime index, map of kid bros
	kidBrosSlice [][]int            // per prime index, slice of kid bros
	resultSets   [][]int            // slice of wanted bro slices
}

func newKidBrothers(setSize, limit int) kidBrothers {
	kbs := kidBrothers{}
	kbs.setSize = setSize
	kbs.limit = limit
	kbs.winnerSum = math.MaxInt32

	kbs.primes, kbs.primeSet = projecteuler.PrimeSet(limit)
	kbs.primeCount = len(kbs.primes)

	kbs.kidBros = make([]map[int]struct{}, kbs.primeCount)
	kbs.kidBrosSlice = make([][]int, kbs.primeCount)

	return kbs
}

func (kbs *kidBrothers) findLowestSetSum() {
	prevResultCount := 0
	set := make([]int, 0, kbs.setSize)

	firstSolution := false
	for i := kbs.setSize - 1; i < kbs.primeCount && !firstSolution; i++ {
		if kbs.primes[i] >= kbs.winnerSum {
			break
		}

		kbs.findKidBros(i)
		if len(kbs.kidBros[i]) < kbs.setSize-1 {
			continue
		}

		set = append(set, i)
		kbs.findBroSet(set)
		set = set[:len(set)-1]

		newResultCount := len(kbs.resultSets)
		for i := prevResultCount; i < newResultCount; i++ {
			sum := kbs.primeSetSum(kbs.resultSets[i])
			if sum < kbs.winnerSum {
				kbs.winnerSum = sum
				firstSolution = true
				break
			}
		}
		prevResultCount = newResultCount
	}
}

func (kbs *kidBrothers) findKidBros(index int) {
	kbs.kidBros[index] = make(map[int]struct{})

	for i := index; i > 0; i-- {
		if kbs.kidBro(index, i-1) {
			kbs.kidBros[index][i-1] = struct{}{}
			kbs.kidBrosSlice[index] = append(kbs.kidBrosSlice[index], i-1)
		}
	}

	return
}

func (kbs kidBrothers) kidBro(x, y int) bool {
	bx := projecteuler.MakeBigIntFromInt(kbs.primes[x])
	by := projecteuler.MakeBigIntFromInt(kbs.primes[y])
	bxCopy := bx.Clone()
	bxCopy.Concatenate(by)
	xy := bxCopy.Int()
	if !kbs.isPrime(xy) {
		return false
	}

	by.Concatenate(bx)
	yx := by.Int()
	if !kbs.isPrime(yx) {
		return false
	}

	return true
}

func (kbs kidBrothers) isPrime(x int64) bool {
	primesLimit := kbs.primes[len(kbs.primes)-1]

	if x <= int64(primesLimit) {
		if _, ok := kbs.primeSet[int(x)]; ok {
			return true
		}
		return false
	}

	root := int64(math.Sqrt(float64(x)))
	for i := 0; i < kbs.primeCount && int64(kbs.primes[i]) <= root; i++ {
		if x%int64(kbs.primes[i]) == 0 {
			return false
		}
	}

	for i := int64(primesLimit) + 2; i <= root; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func (kbs kidBrothers) primeSetSum(set []int) int {
	sum := 0
	for _, x := range set {
		sum += kbs.primes[x]
	}

	return sum
}

func (kbs *kidBrothers) findBroSet(set []int) {
	if len(set) == kbs.setSize {
		newSl := make([]int, kbs.setSize)
		copy(newSl, set)
		kbs.resultSets = append(kbs.resultSets, newSl)

		/*fmt.Printf("%d\t\t[", kbs.primeSetSum(set))
		for i := 0; i < kbs.setSize; i++ {
			fmt.Printf("%d ", kbs.primes[set[i]])
		}
		fmt.Println("]")*/

		return
	}

	wantedKidBros := kbs.setSize - len(set)
	lastInSet := set[len(set)-1]
	existingKidBros := len(kbs.kidBros[lastInSet])
	if existingKidBros < wantedKidBros {
		return
	}

	for i := 0; i < existingKidBros; i++ {
		currKidBro := kbs.kidBrosSlice[lastInSet][i]

		if !kbs.kidBroToPrevious(currKidBro, set) {
			continue
		}

		currKidKidBros := kbs.kidBrosSlice[currKidBro]
		if len(currKidKidBros) < wantedKidBros-1 {
			continue
		}

		foundSetBros := 0
		for j := 0; j < len(currKidKidBros); j++ {
			if kbs.kidBroToPrevious(currKidKidBros[j], set) {
				foundSetBros++
				if foundSetBros == wantedKidBros-1 {
					break
				}
			}
		}

		if foundSetBros == wantedKidBros-1 {
			set = append(set, currKidBro)
			kbs.findBroSet(set)
			set = set[:len(set)-1]
		}
	}
}

func (kbs kidBrothers) kidBroToPrevious(index int, set []int) bool {
	for i := 0; i < len(set)-1; i++ {
		if _, ok := kbs.kidBros[set[i]][index]; !ok {
			return false
		}
	}

	return true
}
