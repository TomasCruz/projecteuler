package main

import (
	"fmt"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 65; Convergents of e

see ./problem065.pdf
*/

func main() {
	projecteuler.Timed(calc, 10000)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	elems := make([]*projecteuler.RootIntElement, 0, (limit+2)/3+1)
	elems = append(elems, &projecteuler.RootIntElement{Head: 2})
	elems = append(elems, &projecteuler.RootIntElement{Head: 1})
	for i := 1; i <= 33; i++ {
		elems = append(elems, &projecteuler.RootIntElement{Head: 2 * i})
	}

	newElems := make([]*projecteuler.RootIntElement, limit)
	newElems[0] = elems[0]
	for i := 1; i < 100; i++ {
		newElems[i] = eElements(elems).eInd(i)
	}

	bif := projecteuler.CalcElements(newElems)
	result = fmt.Sprint(bif.Numerator().DigitSum())
	return
}

type eElements []*projecteuler.RootIntElement

func (e eElements) eInd(x int) *projecteuler.RootIntElement {
	index := 1
	if x%3 == 2 {
		index = x/3 + 2
	}

	return e[index]
}
