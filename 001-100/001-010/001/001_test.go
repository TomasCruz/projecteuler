package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb001(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZWNcdlSYbHyxZ2KtQdWYMGJvzfGNg==", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
