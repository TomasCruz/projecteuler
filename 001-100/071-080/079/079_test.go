package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb079(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BwjToniHqfNo3FjsOPhs9gn7cn3Ys87",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
