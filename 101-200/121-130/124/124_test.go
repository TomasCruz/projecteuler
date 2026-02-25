package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb124(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4mSYzxcVCBn/tGRJg63jYQ63Qx",
		calc,
		100000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
