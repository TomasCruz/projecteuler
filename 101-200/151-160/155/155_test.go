package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb155(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BcnT4/uEP+wpiiCHnYSXl5EkLo1ulM=",
		calc,
		18,
	); err != nil {
		t.Error(err.Error())
	}
}
