package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb050(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDp+JdtpbhHOndS2OeTyLaeHi83wNug==", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
