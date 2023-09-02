package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb041(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BknSo/rFGPv3e5ZE7fuIltjvZ+Kw0Q=",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
