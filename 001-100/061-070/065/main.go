package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 65; Convergents of e

see ./problem065.pdf
*/

func main() {
	var convergentCount int

	if len(os.Args) > 1 {
		convergentCount64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		convergentCount = int(convergentCount64)
	} else {
		convergentCount = 100
	}

	projecteuler.Timed(calc, convergentCount)
}

func calc(args ...interface{}) (result string, err error) {
	convergentCount := args[0].(int)

	elems := make([]*projecteuler.RootIntElement, 0, (convergentCount+2)/3+1)
	elems = append(elems, &projecteuler.RootIntElement{Head: 2})
	elems = append(elems, &projecteuler.RootIntElement{Head: 1})
	for i := 1; i <= (convergentCount-1)/3; i++ {
		elems = append(elems, &projecteuler.RootIntElement{Head: int64(2 * i)})
	}

	newElems := make([]*projecteuler.RootIntElement, convergentCount)
	newElems[0] = elems[0]
	for i := 1; i < convergentCount; i++ {
		newElems[i] = eElements(elems).eInd(i)
	}

	bif := projecteuler.CalcElements(newElems)
	result = fmt.Sprint(digitSum(bif.Numerator()))
	return
}

func digitSum(x *big.Int) int {
	s := x.String()
	r := 0

	for i := 0; i < len(s); i++ {
		r += int(s[i] - '0')
	}

	return r
}

type eElements []*projecteuler.RootIntElement

func (e eElements) eInd(x int) *projecteuler.RootIntElement {
	index := 1
	if x%3 == 2 {
		index = x/3 + 2
	}

	return e[index]
}
