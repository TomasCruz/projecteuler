package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb087(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8rT4juFOD+9ht3NXvgQPdmvGQgGYE=",
		calc,
		50000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
