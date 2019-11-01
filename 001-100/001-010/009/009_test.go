package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb009(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBJeGd9paLnQEL5CtHR195UGvZmBIX69r", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
