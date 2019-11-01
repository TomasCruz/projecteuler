package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb007(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpaKd9tZH/F27q3ml13SBcl0hRyQrA==", calc, 10001); err != nil {
		t.Errorf(err.Error())
	}
}
