package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb146(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8RgkS4jpFaDwt4YjXC9UrjirwZfREOlIMg==",
		calc,
		150000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
