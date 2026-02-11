package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb112(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hoqT4vqF7KvhJlPkyRFL53KL9j0ACg=",
		calc,
		99,
	); err != nil {
		t.Errorf(err.Error())
	}
}
