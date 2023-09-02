package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb034(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x8lS4uzsBXSsdQUpATQfSKErBlw",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
