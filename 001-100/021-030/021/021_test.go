package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb021(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBJeIctld5UrUgE3VcAi3ycISPT1f", calc, 10000); err != nil {
		t.Errorf(err.Error())
	}
}
