package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb142(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8iTorjFIBKhPZSnYfzfcbxj33UFGM=",
		calc,
		4,
	); err != nil {
		t.Errorf(err.Error())
	}
}
