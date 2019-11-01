package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb041(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJCLcttbLeLop9UlnEsUPryEMEzI2ww=", calc); err != nil {
		t.Errorf(err.Error())
	}
}
