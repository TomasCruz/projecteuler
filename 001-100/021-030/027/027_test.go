package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb027(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo6horSojrjATKR4amqM88FScD99nROg==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
