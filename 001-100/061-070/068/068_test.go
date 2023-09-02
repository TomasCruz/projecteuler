package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb068(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8RohSYvpFq7xRdnUzLFFW0PeyzKGTkmqQ4r/GrJF+nE=",
		calc,
		5,
	); err != nil {
		t.Errorf(err.Error())
	}
}
