package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb144(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BomqiyrYYvqAyuDNguYNWTJuw==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
