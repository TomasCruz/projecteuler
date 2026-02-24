package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb127(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hcmSIzjF6M7vjPR30PlHIEgrBiRNQ6D",
		calc,
		120000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
