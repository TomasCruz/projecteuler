package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb062(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA58qEuhypvLAq0wISGFhz1Zo", calc, 5); err != nil {
		t.Errorf(err.Error())
	}
}
