package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb029(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDpeGc8c7r9+90zh3T9l4gL+To6g=", calc, 101); err != nil {
		t.Errorf(err.Error())
	}
}
