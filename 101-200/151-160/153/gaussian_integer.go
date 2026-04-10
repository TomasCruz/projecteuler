package main

type gaussianInteger struct {
	x    [2][2]int64
	norm int64
}

func newGaussianInteger(a, b int64) gaussianInteger {
	return gaussianInteger{
		x: [2][2]int64{
			{a, -b},
			{b, a},
		},
		norm: a*a + b*b,
	}
}

func (g gaussianInteger) add(other gaussianInteger) gaussianInteger {
	r := g.x[0][0] + other.x[0][0]
	i := g.x[1][0] + other.x[1][0]

	return newGaussianInteger(r, i)
}

func (g gaussianInteger) mul(other gaussianInteger) gaussianInteger {
	r := g.x[0][0]*other.x[0][0] + g.x[0][1]*other.x[1][0]
	i := g.x[1][0]*other.x[0][0] + g.x[1][1]*other.x[1][0]

	return newGaussianInteger(r, i)
}

func (g gaussianInteger) div(other gaussianInteger) (gaussianInteger, gaussianInteger) {
	// z1/z2 = (ac+bd)/norm(other)+i*(bc-ad)/norm(other)
	rTimesNormOther := g.x[0][0]*other.x[0][0] + g.x[1][0]*other.x[1][0]
	iTimesNormOther := g.x[1][0]*other.x[0][0] - g.x[0][0]*other.x[1][0]

	r := rTimesNormOther / other.norm
	i := iTimesNormOther / other.norm

	rr := rTimesNormOther - r*other.norm
	ri := iTimesNormOther - i*other.norm

	return newGaussianInteger(r, i), newGaussianInteger(rr, ri)
}
