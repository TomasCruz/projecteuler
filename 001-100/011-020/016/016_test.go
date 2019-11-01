package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb016(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpWIdtzayCoelMt8XGx/2cR8apc=", calc); err != nil {
		t.Errorf(err.Error())
	}
}
