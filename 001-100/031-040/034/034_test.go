package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb034(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5aJc9+wdhTYZs0twm9EpfKkzUsd", calc); err != nil {
		t.Errorf(err.Error())
	}
}
