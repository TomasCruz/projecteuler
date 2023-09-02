package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb053(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x8lTQ8aG7F6sCSZEtbgeyddP/o=",
		calc,
		100,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
