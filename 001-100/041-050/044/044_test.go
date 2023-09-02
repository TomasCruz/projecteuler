package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb044(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8hsqSo3sF6cUINb6D3snZa9/gwqU4R0=",
		calc,
		20000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
