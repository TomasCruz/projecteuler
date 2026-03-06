package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb117(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8iQIviE6L4SNfQyr9AQnn/sof3AgwLd99//Wbw8A==",
		calc,
		50,
	); err != nil {
		t.Errorf(err.Error())
	}
}
