package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb006(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZOPdttbK3R30e6ikjLJ0Op4gq1TrfC8", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
