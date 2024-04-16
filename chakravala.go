package projecteuler

import (
	"fmt"
	"math"
	"math/big"
)

// ChakravalaTriplet holds X, Y, K which are solution to X^2 - N*Y^2 = K
type ChakravalaTriplet struct {
	X *big.Int
	Y *big.Int
	K int64
}

func (c *ChakravalaTriplet) assign(rhs ChakravalaTriplet) {
	c.X.Set(rhs.X)
	c.Y.Set(rhs.Y)
	c.K = rhs.K
}

func (c ChakravalaTriplet) String() string {
	return fmt.Sprintf("[%s, %s, %d]", c.X.String(), c.Y.String(), c.K)
}

// Chakravala returns minimal solution to X^2 - N*Y^2 = 1.
// It uses Chakravala method, commonly attributed to Bhāskara II, a 12th century mathematician.
// His work was building on earlier works by Jayadeva (10th century) and Brahmagupta (7th century).
// Quality and mathematical depth of works by Indian mathematicians regarding numbers and algebra in general
// were only reached by European mathematicians much, much later
func Chakravala(n int) ChakravalaTriplet {
	prev := startingTriplet(n)

	for prev.K != 1 {
		prev.assign(nextTriplet(n, prev))
	}

	return prev
}

func startingTriplet(n int) ChakravalaTriplet {
	var x, k int64

	// x^2 - 61 = k
	xFloat := math.Ceil(math.Sqrt(float64(n)))
	x = int64(xFloat)
	k = x*x - int64(n)

	return ChakravalaTriplet{X: big.NewInt(x), Y: big.NewInt(1), K: k}
}

func nextTriplet(n int, prev ChakravalaTriplet) ChakravalaTriplet {
	m := calcM(n, prev)
	return Samasa(n, m, prev)
}

func Samasa(n int, m *big.Int, prev ChakravalaTriplet) ChakravalaTriplet {
	// samasa of prev (x, y, k) and (m, 1, m^2 - n) gives
	// [math.Abs((xm + ny)/k), math.Abs((my + x)/k), (m^2 - n)/k] with application of Bhaskara's lemma
	// my =(mod k) -x, minimizing m^2 - n
	// iff gcd(y, k) == 1, there is a unique modular multiplicative inverse for b (b^-1)
	// myy^-1 =(mod k) -xy^-1 => m =(mod k) -xy^-1

	kInt, nInt := big.NewInt(prev.K), big.NewInt(int64(n))

	// x
	x, temp, temp2, temp3 := &big.Int{}, &big.Int{}, &big.Int{}, &big.Int{}
	temp.Add(temp2.Mul(m, prev.X), temp3.Mul(nInt, prev.Y))
	x.Div(temp, kInt)
	if x.Sign() == -1 {
		x.Neg(x)
	}

	// y
	y := &big.Int{}
	y.Div(temp2.Add(temp.Mul(m, prev.Y), prev.X), kInt)
	if y.Sign() == -1 {
		y.Neg(y)
	}

	// z
	z := &big.Int{}
	z.Div(temp.Sub(temp2.Mul(m, m), nInt), kInt)

	retValue := ChakravalaTriplet{X: x, Y: y, K: z.Int64()}
	//fmt.Printf("(%d) %s x [%s, 1, %s] = %s\n", n, prev.string(), m.String(), temp.String(), retValue.string())
	return retValue
}

func calcM(n int, prev ChakravalaTriplet) *big.Int {
	kInt := big.NewInt(int64(prev.K))
	a := &big.Int{}
	a.GCD(nil, nil, prev.Y, kInt)
	if a.Int64() == 0 {
		return big.NewInt(0)
	}

	// y^-1
	yInverse := &big.Int{}
	yInverse.ModInverse(prev.Y, kInt)

	// m congruence class, e.g. m =(mod 3) 1
	mCInt := &big.Int{}
	mCInt.Mul(yInverse, prev.X)
	mCInt.Neg(mCInt)
	mCongruent := a.Mod(mCInt, kInt).Int64()

	// number less-equal than root n floor, minus what is neccessary to make it same congruence class as m
	nRootFloor := int64(math.Floor(math.Sqrt(float64(n))))
	rfCongruent := a.Mod(big.NewInt(nRootFloor), kInt).Int64()
	nRootFloor -= rfCongruent - mCongruent

	// number greater-equal than root n floor, plus what is neccessary to make it same congruence class as m
	nRootCeil := nRootFloor + prev.K

	f := math.Abs(float64(nRootFloor*nRootFloor - int64(n)))
	c := math.Abs(float64(nRootCeil*nRootCeil - int64(n)))
	mCongruent = nRootFloor
	if c < f {
		mCongruent = nRootCeil
	}

	return big.NewInt(mCongruent)
}
