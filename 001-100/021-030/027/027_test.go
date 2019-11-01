package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb027(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoGpOHctxbV41m2qlT+pB7WMa+KqEkqw==", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
