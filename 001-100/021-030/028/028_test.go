package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb028(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAZCHcdhbLnRRQg5f3LWBP86eaUlsfg9wHA==", calc, 1001); err != nil {
		t.Errorf(err.Error())
	}
}
