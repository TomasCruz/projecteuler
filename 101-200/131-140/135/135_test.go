package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb135(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xYqQSCCAzdStKSwhQHWY67i2JQ=",
		calc,
		1000000,
	); err != nil {
		t.Error(err.Error())
	}
}
