package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb062(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpSJcNxfJ3FUAYjgEupD3XLCaq/sFFu18z7VGA==", calc, 5); err != nil {
		t.Errorf(err.Error())
	}
}
