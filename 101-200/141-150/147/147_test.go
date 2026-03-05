package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb147(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xskQYrqFa/0SyLo716n+pZmzQfGIqawnA==",
		calc,
		47,
		43,
	); err != nil {
		t.Errorf(err.Error())
	}
}
