package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb051(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpSPc95ZPTzVF1KUZtTRXqys/6jVTg==", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
