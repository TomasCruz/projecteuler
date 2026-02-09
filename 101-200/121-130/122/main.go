package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 122; Efficient Exponentiation
The most naive way of computing n^15 requires fourteen multiplications:
n * n * ... * n = n^15.

But using a "binary" method you can compute it in six multiplications:
n * n = n^2
n^2 * n^2 = n^4
n^4 * n^4 = n^8
n^8 * n^4 = n^12
n^12 * n^2 = n^14
n^14 * n = n^15

However it is yet possible to compute it in only five multiplications:
n * n = n^2
n^2 * n = n^3
n^3 * n^3 = n^6
n^6 * n^6 = n^12
n^12 * n^3 = n^15

We shall define m(k) to be the minimum number of multiplications to compute n^k; for example m(15) = 5.

Find sum[k = 1..200] m(k).
*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 15
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	steps := make([]powerSetSet, limit/4)
	for i := range steps {
		steps[i] = powerSetSet{}
	}
	steps[0][newPowerSet(limit, 1, map[int]struct{}{1: {}})] = struct{}{}
	ms := map[int]int{1: 0}
	missingFirstHalf := map[int]struct{}{}
	for i := 2; i <= limit/2; i++ {
		missingFirstHalf[i] = struct{}{}
	}

	for i := 1; i < len(steps); i++ {
		makeStep(limit, i, steps, ms, missingFirstHalf)
		if len(missingFirstHalf) == 0 {
			break
		}
	}

	mSlice := make([]int, limit+1)
	for i, s := range ms {
		mSlice[i] = s
	}

	powersTwo := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i := limit/2 + 1; i <= limit; i++ {
		if _, present := ms[i]; present {
			continue
		}

		buildMs(i, steps, powersTwo, mSlice)
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		sum += mSlice[i]
	}

	result = strconv.Itoa(sum)
	return
}

func buildMs(x int, steps []powerSetSet, powersTwo, mSlice []int) {
	min := math.MaxInt

	for i := 2; i <= x/2; i++ {
		j := len(powersTwo) - 1
		for ; j > -1; j-- {
			if i*powersTwo[j] <= x {
				break
			}
		}

		x1 := x
		q := x1/i - powersTwo[j]
		x1 -= i * powersTwo[j]
		additionalPowers := 0
		currM := mSlice[i] + j

		for j = j - 1; j > -1 && q > 0; j-- {
			if powersTwo[j] > q {
				continue
			}

			q -= powersTwo[j]
			x1 -= i * powersTwo[j]
			additionalPowers++
		}

		currM += additionalPowers
		if x1 > 0 {
			currM++

			if x1 > 2 && currM < min {
				currM += minSteps(i, x1, steps, mSlice)
			}
		}

		if currM < min {
			min = currM
		}
	}

	mSlice[x] = min
}

func minSteps(d, r int, steps []powerSetSet, mSlice []int) int {
	pssD := []map[int]struct{}{}
	pssR := []map[int]struct{}{}

	for currPS := range steps[mSlice[d]] {
		if currPS[4] == uint64(d) {
			pssD = append(pssD, mapFromPowerSet(currPS))
		}
	}

	for currPS := range steps[mSlice[r]] {
		if currPS[4] == uint64(r) {
			pssR = append(pssR, mapFromPowerSet(currPS))
		}
	}

	min := math.MaxInt
	for _, mr := range pssR {
		if _, present := mr[d]; present {
			return 0
		}

		for _, md := range pssD {
			l := len(mr)
			for xd := range md {
				if _, present := mr[xd]; present {
					l--
				}
			}

			if l < min {
				min = l
			}

			if min == 1 {
				return 1
			}
		}
	}

	return min
}

type powerSet [5]uint64

type powerSetSet map[powerSet]struct{}

func newPowerSet(limit, maxPower int, powersComputed map[int]struct{}) powerSet {
	bs := projecteuler.NewBitset(uint64(limit + 1))

	for k := range powersComputed {
		if k > limit {
			continue
		}

		bs.Set(uint64(k), true)
	}

	sl := []uint64(bs)

	ps := powerSet{}
	copy(ps[:], sl)

	ps[4] = uint64(maxPower)
	return ps
}

func mapFromPowerSet(ps powerSet) map[int]struct{} {
	bs := projecteuler.Bitset(ps[:4])
	return bs.All()
}

func addition(limit int, ps powerSet) []powerSet {
	m := mapFromPowerSet(ps)

	setCardinality := len(m)
	sl := make([]int, setCardinality)
	i := 0
	max := -1
	for k := range m {
		sl[i] = k
		if k > max {
			max = k
		}
		i++
	}

	newX := map[int]struct{}{}
	ret := []powerSet{}
	for i := 0; i < setCardinality; i++ {
		for j := i; j < setCardinality; j++ {
			x := sl[i] + sl[j]
			if _, present := m[x]; present {
				continue
			}

			if x > limit {
				continue
			}

			if _, present := newX[x]; present {
				continue
			}

			newX[x] = struct{}{}

			m[x] = struct{}{}
			newMax := max
			if x > newMax {
				newMax = x
			}
			ret = append(ret, newPowerSet(limit, newMax, m))
			delete(m, x)
		}
	}

	return ret
}

func makeStep(limit, stepOrdinal int, steps []powerSetSet, ms map[int]int, missingFirstHalf map[int]struct{}) {
	for ps := range steps[stepOrdinal-1] {
		nextPowerSets := addition(limit, ps)

		for _, currPS := range nextPowerSets {
			steps[stepOrdinal][currPS] = struct{}{}

			currM := int(currPS[4])
			if _, present := ms[currM]; !present {
				ms[currM] = stepOrdinal
				delete(missingFirstHalf, currM)
				if len(ms) == limit {
					return
				}
			}
		}
	}
}
