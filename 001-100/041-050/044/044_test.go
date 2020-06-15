package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb044(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoApKGctlcLn3Q88/5RZ2aNJx5fbu2PBk=", calc, 20000); err != nil {
		t.Errorf(err.Error())
	}
}
