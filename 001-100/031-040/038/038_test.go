package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb038(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDpWMd95SKHFUm32/fIMcmmAPuhG/FvtLwg==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
