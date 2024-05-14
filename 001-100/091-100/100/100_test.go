package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb100(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BokQIzoFKX3RdbTV20Gj9RpOAZxPN7mcvA5kA==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
