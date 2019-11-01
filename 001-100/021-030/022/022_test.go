package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb022(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoD5GPcdZSLHxSNSouifeJ13LaS+5SXgChFw==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
