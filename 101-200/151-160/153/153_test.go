package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb153(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"",
		calc,
		100000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
