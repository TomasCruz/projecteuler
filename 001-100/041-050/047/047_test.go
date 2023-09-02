package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb047(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hwmSI/pvI3OAkEpzK6z60iAbF0yZA==",
		calc,
		150000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
