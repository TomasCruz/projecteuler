package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb059(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h0rTI/iWY2nBfIOwfn/x03Tjk0Kwg==",
		calc,
		"p059_cipher.txt",
	); err != nil {
		t.Errorf(err.Error())
	}
}
