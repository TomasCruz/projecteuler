package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb055(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZKH0ylmL/QabV/NmwBPwYWNDw==", calc, 10000); err != nil {
		t.Errorf(err.Error())
	}
}
