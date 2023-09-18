package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb080(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x8qQI0i3mwf/+SOyMaGVOV50ZCD",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
