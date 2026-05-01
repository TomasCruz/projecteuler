package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb154(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xgrT4/oE6Lwc/GxU5brUmFaUMe1aacgkg==",
		calc,
		200000,
		"2-12,5-12",
	); err != nil {
		t.Error(err.Error())
	}
}
