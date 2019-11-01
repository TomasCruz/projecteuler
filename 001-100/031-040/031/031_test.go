package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb031(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJWIeN1TFKHW5VosYZK9jp/OMC2i", calc, 200); err != nil {
		t.Errorf(err.Error())
	}
}
