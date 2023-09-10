package projecteuler

import "fmt"

// FibonacciBox is a representation of matrix | A B |
//
//	| C D |,
//
// where A, B are coprime, and B, A, C, D is a generalized Fibonacci sequence,
// i.e. C = B + A, D = C +A. That implies the box can be generated from 2 numbers.
// Every box singularily represents a Pythagorean triplet:
// At = 2AC, Bt = BD, Ct = BC + AD
type FibonacciBox struct {
	A, B, C, D int
}

// PythagoreanTriplet represents triplet for which A^2 + B^2 = C^2
type PythagoreanTriplet struct {
	A, B, C int
}

func (triplet PythagoreanTriplet) Length() int {
	return triplet.A + triplet.B + triplet.C
}

func (triplet PythagoreanTriplet) Multiply(k int) PythagoreanTriplet {
	return PythagoreanTriplet{
		A: k * triplet.A,
		B: k * triplet.B,
		C: k * triplet.C,
	}
}

func (triplet PythagoreanTriplet) String() string {
	return fmt.Sprintf("(%d, %d, %d)", triplet.A, triplet.B, triplet.C)
}

func (box FibonacciBox) Triplet() PythagoreanTriplet {
	return PythagoreanTriplet{
		A: 2 * box.A * box.C,
		B: box.B * box.D,
		C: box.B*box.C + box.A*box.D,
	}
}

type FibonacciBoxTernaryTree struct {
	Root   *FibonacciBox
	Child1 *FibonacciBoxTernaryTree
	Child2 *FibonacciBoxTernaryTree
	Child3 *FibonacciBoxTernaryTree
}

func NewFibonacciBoxTernaryTree(a, b, c, d int) FibonacciBoxTernaryTree {
	return FibonacciBoxTernaryTree{
		Root: &FibonacciBox{
			A: a,
			B: b,
			C: c,
			D: d,
		},
	}
}

func (tree *FibonacciBoxTernaryTree) Generate(limit int) {
	one, two, three := tree.generateRootChildren()

	oneTriplet := one.Triplet()
	if oneTriplet.Length() <= limit {
		childTree := FibonacciBoxTernaryTree{Root: one}
		tree.Child1 = &childTree
		tree.Child1.Generate(limit)
	}

	twoTriplet := two.Triplet()
	if twoTriplet.Length() <= limit {
		childTree := FibonacciBoxTernaryTree{Root: two}
		tree.Child2 = &childTree
		tree.Child2.Generate(limit)
	}

	threeTriplet := three.Triplet()
	if threeTriplet.Length() <= limit {
		childTree := FibonacciBoxTernaryTree{Root: three}
		tree.Child3 = &childTree
		tree.Child3.Generate(limit)
	}
}

func (tree *FibonacciBoxTernaryTree) TripletSlice() []PythagoreanTriplet {
	return tree.extractTriplets()
}

func (tree *FibonacciBoxTernaryTree) extractTriplets() []PythagoreanTriplet {
	triplets := []PythagoreanTriplet{tree.Root.Triplet()}

	if tree.Child1 != nil {
		triplets = append(triplets, tree.Child1.extractTriplets()...)
	}
	if tree.Child2 != nil {
		triplets = append(triplets, tree.Child2.extractTriplets()...)
	}
	if tree.Child3 != nil {
		triplets = append(triplets, tree.Child3.extractTriplets()...)
	}

	return triplets
}

func (tree *FibonacciBoxTernaryTree) generateRootChildren() (one, two, three *FibonacciBox) {
	// box | * b |
	//     | * d | generates boxes
	// | * b |  | b d |  | d b |
	// | d * |, | * * |, | * * |

	o1 := FibonacciBox{
		A: 0,
		B: tree.Root.B,
		C: tree.Root.D,
		D: 0,
	}
	o1.A = o1.C - o1.B
	o1.D = o1.A + o1.C
	one = &o1

	o2 := FibonacciBox{
		A: tree.Root.B,
		B: tree.Root.D,
		C: 0,
		D: 0,
	}
	o2.C = o2.A + o2.B
	o2.D = o2.C + o2.A
	two = &o2

	o3 := FibonacciBox{
		A: tree.Root.D,
		B: tree.Root.B,
		C: 0,
		D: 0,
	}
	o3.C = o3.A + o3.B
	o3.D = o3.C + o3.A
	three = &o3

	return
}
