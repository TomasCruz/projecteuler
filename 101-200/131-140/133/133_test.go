package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb133(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xohTo/tEKf1rDmR6GJq6CJmtybqagEpaA==",
		calc,
		100000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
