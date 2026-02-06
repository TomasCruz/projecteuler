package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb145(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8R8qT4nqhs995fTWCr/XXg53yySwKg==",
		calc,
		9,
	); err != nil {
		t.Errorf(err.Error())
	}
}
