package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb131(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hghQRd+d/lBAAxujEkDFUo0ew==",
		calc,
		6,
	); err != nil {
		t.Errorf(err.Error())
	}
}
