package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb028(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8RkrSYzrF6fxlHquX5h+Q9f28eacgzuzSw==",
		calc,
		1001,
	); err != nil {
		t.Errorf(err.Error())
	}
}
