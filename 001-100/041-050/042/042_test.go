package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb042(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hkg8pujjmmVwQ5dXf7Sq7NWBw==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
