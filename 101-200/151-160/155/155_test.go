package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb155(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"",
		calc,
		18,
	); err != nil {
		t.Error(err.Error())
	}
}
