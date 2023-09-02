package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb023(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x4lQYPtFgNEwTCsSc388sDStAc7xqs=",
		calc,
		28124,
	); err != nil {
		t.Errorf(err.Error())
	}
}
