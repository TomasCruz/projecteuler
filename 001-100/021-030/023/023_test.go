package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb023(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5eJedddL5b7sP8yC9RowYV4mc4ktWs=", calc, 28124); err != nil {
		t.Errorf(err.Error())
	}
}
