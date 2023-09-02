package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb029(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/h4qS5AzRZcBsxd/RI4EoJzNBhk=",
		calc,
		101,
	); err != nil {
		t.Errorf(err.Error())
	}
}
