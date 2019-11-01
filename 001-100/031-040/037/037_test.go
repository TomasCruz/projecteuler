package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb037(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJKGc95dqjOcRRS1CYOu8z6K0TACDA==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
