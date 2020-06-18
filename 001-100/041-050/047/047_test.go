package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb047(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpWKcNtZYQoKtHP78KvWNvqS+YPytQ==", calc, 150000); err != nil {
		t.Errorf(err.Error())
	}
}
