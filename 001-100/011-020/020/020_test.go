package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb020(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8Rsqo7Wl3Qq/aLKn+ixKjsGvGQ==",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
