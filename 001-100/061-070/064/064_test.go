package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb064(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpWMckD55nkuDUFh8xvwGpNijPM=", calc, 10000); err != nil {
		t.Errorf(err.Error())
	}
}
