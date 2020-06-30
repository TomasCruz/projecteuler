package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb056(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDpGMV4dAHyiKJZC90bBjc+qM9g==", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
