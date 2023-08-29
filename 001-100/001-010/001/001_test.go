package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb001(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RwhSY3ijbJn8WcPKGpUFI0gLcO8oA==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
