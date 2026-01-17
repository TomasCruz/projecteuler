package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb102(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R0qdNvhzyliOyLZvp7mQZBAKw==",
		calc,
		1,
	); err != nil {
		t.Errorf(err.Error())
	}
}
