package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb130(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsrSo7pI25rXAk0b/c89yyW5p+kMA==",
		calc,
		25,
	); err != nil {
		t.Errorf(err.Error())
	}
}
