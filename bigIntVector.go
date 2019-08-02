package projecteuler

import (
	"fmt"
	"math/big"
)

var errDim error
var errNegDim error

func init() {
	errDim = fmt.Errorf("dimensions wrong")
	errNegDim = fmt.Errorf("dimensions have to be positive")
}

// BigIntVector struct
type BigIntVector struct {
	vec []*big.Int
}

// Dim returns length of the BigIntVector
func (a *BigIntVector) Dim() int {
	return len(a.vec)
}

func bigIntVectorConstructionHelper(dim int) (result *BigIntVector, err error) {
	if dim < 0 {
		err = errNegDim
		return
	}

	result = &BigIntVector{}
	result.vec = make([]*big.Int, dim)

	return
}

// NewBigIntVector constructs new BigIntVector from slice of int64
func NewBigIntVector(vec []int64) (result *BigIntVector, err error) {
	dim := len(vec)
	if result, err = bigIntVectorConstructionHelper(dim); err != nil {
		return
	}

	for i := 0; i < dim; i++ {
		result.vec[i] = big.NewInt(vec[i])
	}

	return
}

// NewVectorFromBigIntSlice constructs new BigIntVector from slice of *big.Int
func NewVectorFromBigIntSlice(vec []*big.Int) (result *BigIntVector, err error) {
	dim := len(vec)
	if result, err = bigIntVectorConstructionHelper(dim); err != nil {
		return
	}

	for i := 0; i < dim; i++ {
		result.vec[i] = &big.Int{}
		result.vec[i].Set(vec[i])
	}

	return
}

// Clone copies receiver to new BigIntVector
func (a *BigIntVector) Clone() (result *BigIntVector) {
	dim := a.Dim()
	result, _ = bigIntVectorConstructionHelper(dim)

	for i := 0; i < dim; i++ {
		result.vec[i] = &big.Int{}
		result.vec[i].Set(a.vec[i])
	}

	return
}

// MulBigIntVectors multiplies arguments and puts the result in the new big.Int
func MulBigIntVectors(a, b *BigIntVector) (result *big.Int) {
	result = big.NewInt(0)
	for i := 0; i < a.Dim(); i++ {
		currInt := &big.Int{}
		currInt.Mul(a.vec[i], b.vec[i])
		result.Add(result, currInt)
	}

	return
}
