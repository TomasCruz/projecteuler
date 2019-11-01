package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb014(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoD5WJd9ZTdHYTGP1tFa2HFDndSG5ulg==", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
