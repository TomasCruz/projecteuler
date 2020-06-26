package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb053(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5aJdQYgoWCnYnj5MXdyp2lf67E=", calc, 100, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
