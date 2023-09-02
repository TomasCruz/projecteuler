package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb055(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RsrPXEL7ZCxum4/js/lFW0JCg==",
		calc,
		10000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
