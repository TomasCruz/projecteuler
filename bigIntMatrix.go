package projecteuler

import (
	"fmt"
	"math/big"
)

// BigIntMatrix struct
type BigIntMatrix struct {
	m [][]*big.Int
}

// Dim returns length of the BigIntMatrix
func (a *BigIntMatrix) Dim() (x, y int) {
	return len(a.m[0]), len(a.m)
}

// At returns element at (y,x)
func (a *BigIntMatrix) At(y, x int) (result *big.Int, err error) {
	if x < 0 || y < 0 {
		err = errNegDim
		return
	}

	dx, dy := a.Dim()
	if x > dx || y > dy {
		err = errDim
		return
	}

	result = &big.Int{}
	result.Set(a.m[y][x])
	return
}

// Print prints the BigIntMatrix
func (a *BigIntMatrix) Print() {
	x, y := len(a.m[0]), len(a.m)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Print(a.m[i][j].String())
		}
		fmt.Println()
	}
}

func bigIntMatrixConstructionHelper(x, y int) (result *BigIntMatrix, err error) {
	if x <= 0 || y <= 0 {
		err = errNegDim
		return
	}

	result = &BigIntMatrix{}
	result.m = make([][]*big.Int, y)

	for i := 0; i < y; i++ {
		result.m[i] = make([]*big.Int, x)
	}

	return
}

// NewUnitBigIntMatrix constructs new unit BigIntMatrix
func NewUnitBigIntMatrix(x int) (result *BigIntMatrix, err error) {
	if result, err = bigIntMatrixConstructionHelper(x, x); err != nil {
		return
	}

	for i := 0; i < x; i++ {
		result.m[i][i] = big.NewInt(int64(1))
	}

	return
}

// NewBigIntMatrix constructs new BigIntMatrix
func NewBigIntMatrix(x, y int, m []int64) (result *BigIntMatrix, err error) {
	if len(m) != x*y {
		err = errDim
		return
	}

	if result, err = bigIntMatrixConstructionHelper(x, y); err != nil {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			result.m[i][j] = big.NewInt(m[i*x+j])
		}
	}

	return
}

// Clone copies receiver to new BigIntMatrix
func (a *BigIntMatrix) Clone() (result *BigIntMatrix) {
	x, y := a.Dim()
	result, _ = bigIntMatrixConstructionHelper(x, y)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			result.m[i][j] = &big.Int{}
			result.m[i][j].Set(a.m[i][j])
		}
	}

	return
}

// MulBigIntMatrix multiplies arguments and puts the result in the new BigIntMatrix
func MulBigIntMatrix(a, b *BigIntMatrix) (result *BigIntMatrix) {
	_, ay := a.Dim()
	bx, by := b.Dim()

	result, _ = bigIntMatrixConstructionHelper(bx, ay)

	for i := 0; i < ay; i++ {
		for j := 0; j < bx; j++ {
			bSlice := make([]*big.Int, by)
			for k := 0; k < by; k++ {
				bSlice[k] = b.m[k][j]
			}

			aVec, _ := NewVectorFromBigIntSlice(a.m[i])
			bVec, _ := NewVectorFromBigIntSlice(bSlice)
			result.m[i][j] = MulBigIntVectors(aVec, bVec)
		}
	}

	return
}
