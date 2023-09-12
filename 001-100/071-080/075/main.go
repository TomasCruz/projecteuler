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

	count := 0
	for i := range lengths {
		if lengths[i] == 1 {
			count++
		}
	}

	result = strconv.Itoa(count)
	return
}
