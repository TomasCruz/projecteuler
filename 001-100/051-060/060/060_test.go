package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb060(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZCOc9xz90oMEKlb/mO6khxDroPA", calc, 5, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
