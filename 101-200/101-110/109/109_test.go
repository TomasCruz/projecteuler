package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb109(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BcjQIlsRHorZf4GQKTHywmWTByr",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
