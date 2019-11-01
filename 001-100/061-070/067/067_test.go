package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb067(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJSJcxQQ3lQCI14hZY3bpQweXmo=", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
