package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb057(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpON0Th8HL/FWFKwQI04vloj9g==", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
