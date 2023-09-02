package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb061(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RckQI8M+82k+JxdDcwF4DOEtJN4",
		calc,
		6,
		projecteuler.ConcurrentKind,
	); err != nil {

		t.Errorf(err.Error())
	}
}
