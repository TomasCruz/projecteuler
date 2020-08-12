package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb066(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdoAZCPGSshz7O2dN0lmKrWn/HNyg==",
		calc, 1000, chakravalaSolver{}); err != nil {

		t.Errorf(err.Error())
	}
}
