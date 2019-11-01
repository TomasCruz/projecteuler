package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb020(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAZKGS27angPWbK7YkL8eo0Q/DQ==", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
