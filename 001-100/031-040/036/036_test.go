package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb036(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoD5GMcdddACqagcS48N3Twcz/FFaNQg==", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
