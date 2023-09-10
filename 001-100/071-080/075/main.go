package main

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 75; Singular Integer Right Triangles
It turns out that 12cm is the smallest length of wire that can be bent to form an integer sided right angle triangle
in exactly one way, but there are many more examples.
12cm: (3,4,5)
24cm: (6,8,10)
30cm: (5,12,13)
36cm: (9,12,15)
40cm: (8,15,17)
48cm: (12,16,20)
In contrast, some lengths of wire, like 20, cannot be bent to form an integer sided right angle triangle,
and other lengths allow more than one solution to be found; for example, using 120cm it is possible
to form exactly three different integer sided right angle triangles.
120cm: (30,40,50), (20,48,52), (24,45,51)
Given that L is the length of the wire, for how many values of L <= 1500000 can exactly one
integer sided right angle triangle be formed?
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
		limit = 1500000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// primes := projecteuler.Primes(limit, nil)
	tree := projecteuler.NewFibonacciBoxTernaryTree(1, 1, 2, 3)
	tree.Generate(limit)
	triplets := tree.TripletSlice()
	tripletsLength := len(triplets)

	tripletLengthSlice := make([]int, tripletsLength)
	for i := 0; i < tripletsLength; i++ {
		tripletLengthSlice[i] = triplets[i].Length()
	}
	sort.Ints(tripletLengthSlice)

	lengths := make([]int, limit+1)
	for i := range tripletLengthSlice {
		for k := 1; ; k++ {
			x := k * tripletLengthSlice[i]
			if x > limit {
				break
			}
			lengths[x]++
		}
	}

	// k := 0
	// fmt.Printf("%d", tripletLengthSlice[0])
	// for i := 1; i < tripletsLength; i++ {
	// 	fmt.Printf(", %d", tripletLengthSlice[i])
	// 	if k == 20 {
	// 		fmt.Println()
	// 		k = 0
	// 	}
	// 	k++
	// }
	// fmt.Println()
	// fmt.Println()

	count := 0
	for i := range lengths {
		if lengths[i] == 1 {
			count++
		}
	}

	// for lengthOfWire := 12; lengthOfWire <= limit; lengthOfWire++ {
	// 	// find smallest tripletLengthSlice index greater or equal to lengthOfWire
	// 	currLimit := binarySearch(lengthOfWire, tripletLengthSlice, 0, tripletsLength-1)
	// 	// fmt.Printf("binarySearch %d: %d\n", lengthOfWire, currLimit)

	// 	wg := sync.WaitGroup{}
	// 	threadPerCoreGuard := make(chan int, runtime.GOMAXPROCS(runtime.NumCPU()))
	// 	solvingResult := make([]int, currLimit)
	// 	for i := 0; i < currLimit; i++ {
	// 		threadPerCoreGuard <- 1
	// 		wg.Add(1)
	// 		go func(i, x, y int) {
	// 			m := x % y
	// 			if m == 0 {
	// 				solvingResult[i] = 1
	// 			}

	// 			<-threadPerCoreGuard
	// 			wg.Done()
	// 		}(i, lengthOfWire, tripletLengthSlice[i])
	// 	}
	// 	wg.Wait()

	// 	currSolutionCount := 0
	// 	for _, x := range solvingResult {
	// 		currSolutionCount += x
	// 		if currSolutionCount > 1 {
	// 			break
	// 		}
	// 	}
	// 	if currSolutionCount == 1 {
	// 		count++
	// 	}

	// oneFound := false
	// twoFound := false
	// for i := 0; i < tripletsLength; i++ {
	// 	currLength := tripletLengthSlice[i]
	// 	if currLength > lengthOfWire {
	// 		break
	// 	}

	// 		// // sub-puzzles
	// 		// var subPuzzles []*Sudoku
	// 		// uns := s.firstUnsolved()
	// 		// wg := sync.WaitGroup{}
	// 		// threadPerCoreGuard := make(chan int, runtime.GOMAXPROCS(runtime.NumCPU()))
	// 		// solvingResult := make(chan int, length)

	// 		// for x := range s.marks[uns.r][uns.c] {
	// 		// 	s1 := &Sudoku{}
	// 		// 	s1.copy(s)
	// 		// 	s1.solveCell(uns, x)
	// 		// 	subPuzzles = append(subPuzzles, s1)

	// 		// 	threadPerCoreGuard <- 1
	// 		// 	wg.Add(1)
	// 		// 	go func(sud *Sudoku) {
	// 		// 		sud.Solve()
	// 		// 		<-threadPerCoreGuard
	// 		// 		wg.Done()
	// 		// 	}(s1)
	// 		// }
	// 		// wg.Wait()

	// 		// for s := range solvingResult {
	// 		// 	resInt += s
	// 		// }

	// 	if lengthOfWire%currLength == 0 {
	// 		if oneFound {
	// 			twoFound = true
	// 			break
	// 		} else {
	// 			oneFound = true
	// 		}
	// 	}
	// }

	// if oneFound && !twoFound {
	// 	count++
	// }
	// }

	result = strconv.Itoa(count)
	return
}

func binarySearch(x int, slice []int, firstIndex, lastIndex int) int {
	if lastIndex-firstIndex <= 2 {
		return lastIndex
	}

	middleIndex := (firstIndex + lastIndex) / 2
	if x == slice[middleIndex] {
		return middleIndex + 1
	} else if x < slice[middleIndex] {
		return binarySearch(x, slice, firstIndex, middleIndex)
	}

	if middleIndex < lastIndex && slice[middleIndex+1] > x {
		return middleIndex + 1
	}

	return binarySearch(x, slice, middleIndex+1, lastIndex)
}
