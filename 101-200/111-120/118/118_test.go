package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb118(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xskQIu5A93SguLFZ0H3ns8lqVhg",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
