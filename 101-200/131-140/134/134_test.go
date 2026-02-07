package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb134(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hckSYjuFaH2R9LWz7FGXzSi9i4qTKBK2kjB3P8PsJXz",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
