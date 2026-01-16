package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb101(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BgiT43rFqP1Q9e3OuRI5qvZyY8x0RVBUJQu",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
