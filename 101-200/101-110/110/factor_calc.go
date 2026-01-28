package main

import (
	"sort"

	"github.com/TomasCruz/projecteuler"
)

type factorType struct {
	n          int64
	divSquareN int
	factors    []int
}

type ftSlice []factorType

func (a ftSlice) Len() int           { return len(a) }
func (a ftSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ftSlice) Less(i, j int) bool { return a[i].n < a[j].n }

type factorCalc struct {
	limit       int
	primePowers [][]int64
}

func newFC(limit int) factorCalc {
	primes := projecteuler.Primes(100, nil)
	primePowers := buidPrimePowers(primes)
	return factorCalc{limit: limit, primePowers: primePowers}
}

func (fc factorCalc) calcSubsetFactorTypes(n int64, maxFactors, dontExceedFactors []int) ftSlice {
	ret := make(ftSlice, 0)
	factors := make([]int, len(dontExceedFactors))

	ft := factorType{
		n:          int64(1),
		divSquareN: 1,
		factors:    factors,
	}
	ret = fc.calcSubsetFactorTypesRec(n, maxFactors, dontExceedFactors, ft, 0, ret)

	sort.Sort(ret)

	return ret
}

func (fc factorCalc) calcSubsetFactorTypesRec(n int64, maxFactors, dontExceedFactors []int, ft factorType, from int, ret ftSlice) ftSlice {
	l := len(dontExceedFactors)
	if from == l {
		if ft.divSquareN >= 2*fc.limit {
			newFactors := make([]int, l)
			copy(newFactors, ft.factors)
			ft.factors = newFactors
			ret = append(ret, ft)
		}
		return ret
	}

	for i := 0; i <= dontExceedFactors[from]; i++ {
		newN := ft.n * fc.primePowers[from][i]
		if newN > n {
			break
		}

		oldN := ft.n
		oldDivSqN := ft.divSquareN
		oldFactorsFrom := ft.factors[from]
		ft.n = newN
		ft.divSquareN = ft.divSquareN * (2*i + 1)
		ft.factors[from] = i
		ret = fc.calcSubsetFactorTypesRec(n, maxFactors, dontExceedFactors, ft, from+1, ret)
		ft.n = oldN
		ft.divSquareN = oldDivSqN
		ft.factors[from] = oldFactorsFrom
	}

	return ret
}

func (fc factorCalc) findDivisorsLesserThanRootCount(ft factorType) int {
	divSqN := 1
	factorsSqN := make([]int, len(ft.factors))
	for i, x := range ft.factors {
		divSqN *= (2*x + 1)
		factorsSqN[i] = 2 * x
	}
	sqFT := factorType{
		divSquareN: divSqN,
		factors:    factorsSqN,
	}

	ret := fc.calcSubsets(ft, sqFT)
	return len(ret)
}

func (fc factorCalc) calcSubsets(nFT, sqFT factorType) ftSlice {
	ret := make(ftSlice, 0)
	ft := factorType{
		n:       int64(1),
		factors: make([]int, len(sqFT.factors)),
	}
	ret = fc.calcSubsetsRec(nFT, sqFT, ft, 0, ret)

	return ret
}

func (fc factorCalc) calcSubsetsRec(nFT, sqFT, ft factorType, from int, ret ftSlice) ftSlice {
	l := len(sqFT.factors)
	if from == l {
		newFactors := make([]int, l)
		copy(newFactors, ft.factors)
		ft.factors = newFactors
		return append(ret, ft)
	}

	for i := 0; i <= sqFT.factors[from]; i++ {
		newN := ft.n * fc.primePowers[from][i]
		if newN > nFT.n {
			break
		}

		newFactors := make([]int, l)
		copy(newFactors, ft.factors)
		newFactors[from] = i

		newFT := factorType{
			n:       newN,
			factors: newFactors,
		}

		oldN := ft.n
		oldFactorsFrom := ft.factors[from]
		ft.n = newN
		ft.factors[from] = i
		ret = fc.calcSubsetsRec(nFT, sqFT, newFT, from+1, ret)
		ft.n = oldN
		ft.factors[from] = oldFactorsFrom
	}

	return ret
}
