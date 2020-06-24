package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb048(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDpePcNdeKHNQB3qA8gNeRlNy4ZkkxQISjYo=", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
