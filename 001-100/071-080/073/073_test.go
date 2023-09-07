package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb073(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8B0rTYjtFUmHe0yFxc6HbyLJUtnAQyk=",
		calc,
		12000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
