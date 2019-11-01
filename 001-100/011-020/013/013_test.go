package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb013(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoApONd9xdKHZTB7M6S0w2rb56D0GQ+BIRRu4=", calc, 10); err != nil {
		t.Errorf(err.Error())
	}
}
