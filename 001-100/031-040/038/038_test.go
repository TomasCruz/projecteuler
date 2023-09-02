package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb038(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/hwgT4riEaL0kpzLyfeFz6S9cRPb7S9k1g==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
