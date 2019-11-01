package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb003(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAZ6Ld6Smn81NRjmWr3sv/TfbP6Q=",
		calc, int64(600851475143)); err != nil {

		t.Errorf(err.Error())
	}
}
