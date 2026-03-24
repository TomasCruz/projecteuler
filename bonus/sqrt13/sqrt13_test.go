package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProbSqrt13(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"",
		calc,
		13,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
