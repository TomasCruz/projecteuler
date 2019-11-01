package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb024(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZGGc9ZbK3BWBz8cT2ENPHS24+GgeHFHrO8=",
		calc, byte(10), 1000000); err != nil {

		t.Errorf(err.Error())
	}
}
