package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb059(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpSHdNtSfUSNZxdBs53juTNWnNclUg==", calc, "p059_cipher.txt"); err != nil {
		t.Errorf(err.Error())
	}
}
