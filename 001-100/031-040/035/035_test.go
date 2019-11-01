package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb035(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoApMqPv0QDEE+eeXBEQacgYdt", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
