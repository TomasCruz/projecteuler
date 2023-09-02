package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb022(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xgjSYLiFa/y2m719lp/vcVtXOg8ePdh/g==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
