package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb108(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hciSYPqrxSCKuzoyaP54ghm0BiKsA==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
