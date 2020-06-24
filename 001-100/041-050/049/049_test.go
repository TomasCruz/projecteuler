package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb049(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZ+IedlYJ31ZAYLqhJCxYFfed65FrSG0P63McQ==", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
