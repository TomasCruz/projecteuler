package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb052(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsgQI7tCIpT9x4YSHJnIPkDkrcZRw==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
