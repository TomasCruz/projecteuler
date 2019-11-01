package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb005(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZWMd9ZYK3JQ2WsS+RE0r9QsiWRuNSN0Sg==", calc, 20); err != nil {
		t.Errorf(err.Error())
	}
}
