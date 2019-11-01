package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb025(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5GGcgYnXMMcUBQEBAR2dbjgLcA=", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
