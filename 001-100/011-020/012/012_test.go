package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb012(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJCLd9lfLnRrlHBJao66cclrclRuqIMN", calc, 500); err != nil {
		t.Errorf(err.Error())
	}
}
