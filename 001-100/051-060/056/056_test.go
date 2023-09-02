package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb056(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/hggRh43ZbSUEEOuwVhSXqzMTg==",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
