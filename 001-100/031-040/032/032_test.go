package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb032(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xogSoPKUmPB1fIsATCF9RpUSjs4",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
