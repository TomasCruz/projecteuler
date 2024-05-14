package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb099(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8B8rotfFFq8YHT00q6KLnOqIWg==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
