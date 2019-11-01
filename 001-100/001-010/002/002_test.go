package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb002(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5CPc9hZLHkX0KPIj/7e3zr+FPEeB4A=",
		calc, int64(4000000)); err != nil {

		t.Errorf(err.Error())
	}
}
