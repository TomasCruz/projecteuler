package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb004(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/h8kTovjtur8+2hJh3nLP0BdFVWY4g==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
