package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb003(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8RcnTxhJ2j+Tr79fbyrFYvdcW6Q=",
		calc,
		int64(600851475143),
	); err != nil {

		t.Error(err.Error())
	}
}
