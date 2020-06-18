package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb046(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoApGJd0eBMydYrwaMuHLLN/ffGSc=", calc, 10000); err != nil {
		t.Errorf(err.Error())
	}
}
