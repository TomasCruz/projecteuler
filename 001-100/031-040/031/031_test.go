package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb031(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BwkQIkH8wTBB7Ek/ZtDW5rgS7Oe",
		calc,
		200,
	); err != nil {
		t.Errorf(err.Error())
	}
}
